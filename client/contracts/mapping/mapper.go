package mapping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/assets"
)

type SourceMapper struct {
	contracts             []*Contract
	sourceMapInstructions []SourceMapInstruction
	files                 []string
	contractsPath         string
	useBinding            bool
}

type Contract struct {
	instructions          []*Instruction
	sourceMapInstructions []*SourceMapInstruction
}

type Config struct {
	ContractsPath string
	UseBinding    bool
}

func (c *Contract) GetInstructionByPc(pc uint64) (*Instruction, *SourceMapInstruction, error) {
	insInMap, err := c.getByteIndexByPc(pc)
	if err != nil {
		return nil, nil, err
	}

	ins := c.instructions[insInMap]
	smIns := c.sourceMapInstructions[insInMap]

	return ins, smIns, nil
}

func (c *Contract) getByteIndexByPc(pc uint64) (int, error) {
	for i, ins := range c.instructions {
		if ins.byteCodePosition >= int(pc) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("instruction out of bounds")
}

type JSONSourceMap struct {
	Contracts  map[string]contract `json:"contracts"`
	Version    string              `json:"version"`
	SourceList []string            `json:"sourceList"`
}

type contract struct {
	BinRuntime    string `json:"bin-runtime"`
	SrcMapRuntime string `json:"srcmap-runtime"`
}

func NewFromCombinedRuntimeFile(filePath string, config *Config) (*SourceMapper, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading source map: %s", err)
	}

	return NewFromCombinedRuntime(fileContent, config)
}

func NewFromCombinedRuntime(fileContent []byte, config *Config) (*SourceMapper, error) {
	if config == nil {
		config = &Config{}
	}

	sourceMap := JSONSourceMap{}

	decoder := json.NewDecoder(bytes.NewReader(fileContent))
	err := decoder.Decode(&sourceMap)
	if err != nil {
		return nil, fmt.Errorf("error decoding source map: %s", err)
	}

	contracts, err := parseContracts(sourceMap)
	if err != nil {
		return nil, fmt.Errorf("error parsing contract data: %s", err)
	}

	return &SourceMapper{
		files:         parseFiles(sourceMap),
		contracts:     contracts,
		contractsPath: config.ContractsPath,
		useBinding:    config.UseBinding, // Uses the bindings from contracts/bindings/assets intead of file path.
	}, nil
}

func (sm *SourceMapper) GetFileByIndex(index int) (string, error) {
	if len(sm.files) <= index {
		return "", fmt.Errorf("invalid index for file")
	}

	return sm.files[index], nil
}

func (sm *SourceMapper) GetContractByIndex(index int) (*Contract, error) {
	if len(sm.contracts) <= index {
		return nil, fmt.Errorf("invalid index for contract")
	}

	return sm.contracts[index], nil
}

func (sm *SourceMapper) GetSolidityLineByPc(pc uint64) (string, error) {
	contract, err := sm.GetContractByIndex(0)
	if err != nil {
		return "", fmt.Errorf("error getting contract bt pc: %s", err)
	}

	_, insMap, err := contract.GetInstructionByPc(pc)
	if err != nil {
		return "", fmt.Errorf("error getting instruction by pc: %s", err)
	}

	fileName, err := sm.GetFileByIndex(insMap.fileIndex)
	if err != nil {
		return "", fmt.Errorf("error getting contract file: %s", err)
	}

	var fileContents []byte
	if sm.useBinding {
		fileContents, err = assets.Asset(fileName)
		if err != nil {
			return "", fmt.Errorf("error getting contract data from bindings: %s", err)
		}
	} else {
		contractsPath := filepath.Join(sm.contractsPath, fileName)
		fileContents, err = ioutil.ReadFile(contractsPath)
		if err != nil {
			return "", fmt.Errorf("error getting contents of contract file: %s", err)
		}
	}

	return string(fileContents[insMap.byteOffsetStart : insMap.byteOffsetStart+insMap.sourceRangeLength]), nil
}

func parseFiles(jsm JSONSourceMap) []string {
	var files []string

	for _, file := range jsm.SourceList {
		files = append(files, file)
	}

	return files
}

func parseContracts(jsm JSONSourceMap) ([]*Contract, error) {
	var contracts []*Contract
	var keys []string

	for k := range jsm.Contracts {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Force ordering by string to avoid randomness in maps.

	for _, c := range keys {
		theContract := jsm.Contracts[c]
		smi, err := ParseSourceMap(theContract.SrcMapRuntime)
		if err != nil {
			return nil, fmt.Errorf("error parsing source map: %s", err)
		}

		ins, err := ParseByteCode(common.Hex2Bytes(theContract.BinRuntime))
		if err != nil {
			return nil, fmt.Errorf("error parsing bytecode: %s", err)
		}

		contracts = append(contracts, &Contract{
			sourceMapInstructions: smi,
			instructions:          ins,
		})
	}

	return contracts, nil
}
