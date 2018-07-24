pragma solidity 0.4.24;

import "openzeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "../consensus/mgr/ValidatorMgr.sol";

contract OracleMgr is Pausable {

    uint public baseDeposit;       
    uint public maxNumOracles;
    uint public freezePeriod;
    uint public syncFrequency;
    uint public updatePeriod;
    uint public averagePrice = 0;
    ValidatorMgr validatorMgr;
    bytes4 sig = bytes4(keccak256("isSuperNode(address)"));

    struct Deposit {
        uint amount;
        uint availableAt;
    }

    struct Oracle {
        uint index;
        bool isOracle;
        bool hasSubmittedPrice;
        Deposit[] deposits; 
    }
    
    mapping (address => Oracle) private oracleRegistry;
    
    // oraclePool contains the oracle identity ordered by the biggest deposit to
    // the smallest deposit.
    address[] private oraclePool;

    address[] private submissions;

    modifier onlyWithMinDeposit {
        require(msg.value >= getMinimumDeposit());
        _;
    }

    modifier onlyOnce {
        require(!oracleRegistry[msg.sender].hasSubmittedPrice);
        _;
    }

    modifier onlyOracle {
        require(isOracle(msg.sender));
        _;
    }

    modifier onlyNewCandidate {
        require(!isOracle(msg.sender));
        _;
    }

    modifier onlySuperNode {
        require(validatorMgr.isSuperNode(msg.sender));
        _;
    }

    modifier onlyValidPrice(uint _price) {
        require(_price > 0);
        _;
    }

    function OracleMgr(
        uint _baseDeposit,
        uint _maxNumOracles,
        uint _freezePeriod,
        uint _syncFrequency,
        uint _updatePeriod,
        address _validatorMgrAddr) 
    public {
        require(_maxNumOracles > 0);
        require(_syncFrequency >= 0);

        // sync enabled
        if (_syncFrequency > 0) {
            require(_updatePeriod > 0 && _updatePeriod <= _syncFrequency);
        }
        
        baseDeposit = _baseDeposit;
        maxNumOracles = _maxNumOracles;
        freezePeriod = _freezePeriod * 1 days;
        syncFrequency = _syncFrequency;
        updatePeriod = _updatePeriod;
        validatorMgr = ValidatorMgr(_validatorMgrAddr);
    }

    function isOracle(address identity) public view returns (bool isIndeed) {
        return oracleRegistry[identity].isOracle;
    }

    function _hasAvailability() public view returns (bool available) {
        return (maxNumOracles - oraclePool.length) > 0;
    }

    function _deleteOracle(address identity) private {
        Oracle oracle = oracleRegistry[identity];
        for (uint index = oracle.index; index < oraclePool.length - 1; index++) {
            oraclePool[index] = oraclePool[index + 1];
        }
        oraclePool.length--;

        oracle.isOracle = false;
        oracle.deposits[oracle.deposits.length - 1].availableAt = now + freezePeriod;
    }

    // _deleteSmallestBidder removes the oracle with the smallest deposit
    function _deleteSmallestBidder() private {
        _deleteOracle(oraclePool[oraclePool.length - 1]);
    }

    function _insertOracle(address identity, uint deposit) private {
        Oracle oracle = oracleRegistry[identity];
        oracle.index = oraclePool.push(identity) - 1;
        oracle.isOracle = true;
        oracle.deposits.push(Deposit({amount: deposit, availableAt: 0}));

        for (uint index = oracle.index; index > 0; index--) {
            Oracle target = oracleRegistry[oraclePool[index - 1]];
            Deposit collateral = target.deposits[target.deposits.length - 1];
            if (deposit <= collateral.amount) {
                break;
            }
            oraclePool[index] = oraclePool[index - 1];
            oraclePool[index - 1] = identity; 
            // update indexes
            target.index = index;
            oracle.index = index - 1;
        }
    }

    function getOracleCount() public view returns (uint count) {
        return oraclePool.length;
    }

    function getOracleAtIndex(uint index) public view returns (address code, uint deposit) {
        code = oraclePool[index];
        Oracle oracle = oracleRegistry[code];
        deposit = oracle.deposits[oracle.deposits.length - 1].amount;
    }

    function getNumSubmissions() public view returns (uint count) {
        return submissions.length;
    }

    function getSubmissionAtIndex(uint index) public view returns (address identity) {
        return submissions[index];
    }

    // getMinimumDeposit returns the base deposit if there are positions available or
    // the current smallest deposit required if there aren't positions available.
    function getMinimumDeposit() public view returns (uint deposit) {
        // there are positions for validator available
        if (_hasAvailability()) {
            return baseDeposit;
        } else {
            Oracle smallestBidder = oracleRegistry[oraclePool[oraclePool.length - 1]];               
            return smallestBidder.deposits[smallestBidder.deposits.length - 1].amount + 1;
        }
    }
    
    function getDepositCount() public view returns (uint count) {
        return oracleRegistry[msg.sender].deposits.length; 
    }

    function getDepositAtIndex(uint index) public view returns (uint amount, uint availableAt) {
        Deposit deposit = oracleRegistry[msg.sender].deposits[index];
        return (deposit.amount, deposit.availableAt);
    }

    // registerOracle registers a new candidate as oracle
    function registerOracle() public payable whenNotPaused onlyNewCandidate onlyWithMinDeposit onlySuperNode {
        if (!_hasAvailability()) {
            _deleteSmallestBidder();
        }
        _insertOracle(msg.sender, msg.value);
    }

    // deregisterOracle deregisters the msg sender from the oracle set
    function deregisterOracle() public whenNotPaused onlyOracle {
        _deleteOracle(msg.sender);
    }

    function _removeDeposits(address identity, uint index) private {
        if (index == 0) return;

        Oracle oracle = oracleRegistry[identity];
        uint lo = 0;
        uint hi = index;
        while (hi < oracle.deposits.length) {
            oracle.deposits[lo] = oracle.deposits[hi];
            lo++;
            hi++;
        }
        oracle.deposits.length = lo;
    }

    // releaseDeposits transfers locked deposit(s) back the user account if they
    // are past the freeze period
    function releaseDeposits() public whenNotPaused {
        uint refund = 0;
        uint i = 0;
        Deposit[] deposits = oracleRegistry[msg.sender].deposits;
        
        for (; i < deposits.length && deposits[i].availableAt != 0; i++) {
            if (now < deposits[i].availableAt) {
                break;
            }
            refund += deposits[i].amount;
        }
        
        _removeDeposits(msg.sender, i);

        if (refund > 0) {
            msg.sender.transfer(refund);
        }
    }

    function submitPrice(uint _price) public whenNotPaused onlyOracle onlyOnce onlyValidPrice(_price) {
        averagePrice = _price;
        oracleRegistry[msg.sender].hasSubmittedPrice = true;
        submissions.push(msg.sender);
    }
}