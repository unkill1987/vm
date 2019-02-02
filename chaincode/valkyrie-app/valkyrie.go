package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}


type Trade struct {
	Os  	Osform `json:"os"`
	Lcr 	Lcrform `json:"lcr"`
	Lc	Lcform `json:"lc"`
	Sr 	Srform `json:"sr"`
	Bl	Blform `json:"bl"`
	Ci	Ciform `json:"ci"`
	Do	Doform `json:"do"`
}

type Osform struct {
	From  	string `json:"from"`
	To 	string `json:to`
	Doc_hash	string `json:"doc_hash"`
}

type Lcrform struct {
	From  	string `json:"from"`
	To 	string `json:to`
	Doc_hash	string `json:"doc_hash"`
}

type Lcform struct {
	From  	string `json:"from"`
	To 	string `json:to`
	Doc_hash	string `json:"doc_hash"`
}

type Srform struct {
	From  	string `json:"from"`
	To 	string `json:to`
	Doc_hash	string `json:"doc_hash"`
}

type Blform struct {
	From  	string `json:"from"`
	To 	string `json:to`
	Doc_hash	string `json:"doc_hash"`
}

type Ciform struct {
	From  	string `json:"from"`
	To 	string `json:to`
	Doc_hash	string `json:"doc_hash"`
}

type Doform struct {
	From  	string `json:"from"`
	To 	string `json:to`
	Doc_hash	string `json:"doc_hash"`
}




func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()

	if function =="queryTrade" {
		return s.queryTrade(APIstub, args) 
	}else if function == "recordOS"{
		return s.recordOS(APIstub, args) 
	}else if function == "recordLCR"{
		return s.recordLCR(APIstub, args) 
	}else if function == "recordLC"{
		return s.recordLC(APIstub, args) 
	}else if function == "recordSR"{
		return s.recordSR(APIstub, args) 
	}else if function == "recordBL"{
		return s.recordBL(APIstub, args)  
	}else if function == "recordCI"{
		return s.recordCI(APIstub, args)   
	}else if function == "recordDO"{
		return s.recordDO(APIstub, args)  
	}else if function == "keyHistory"{
		return s.keyHistory(APIstub, args)
	}
	return shim.Error("Invalid SmartContract function name.")
}

func (s *SmartContract) keyHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) !=1{
		return shim.Error("Incorrect number of arguments.Expecting 1")
	}
	contract_id := args[0]
	history, err := APIstub.GetHistoryForKey(contract_id)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer history.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for history.HasNext() {
		response, err := history.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
	
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Value\":")
		buffer.WriteString(string(response.Value))
		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("Success", buffer.String())
	
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryTrade(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1{
		return shim.Error("Incorrect number of arguments.Expecting 1")
	}
	
	tradeAsBytes, _ :=APIstub.GetState(args[0])
	
	if tradeAsBytes == nil {
		return shim.Error("Could not found Contract_ID")
	}
	return shim.Success(tradeAsBytes)
}

func (s *SmartContract) recordOS(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	
	//  0-contract_id  1-exporter  2-importer  3-os_hash
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	fmt.Println("start init contract")
	
	contract_id:= args[0]
	exporter := strings.ToLower(args[1])
	importer := strings.ToLower(args[2])
	os_hash := strings.ToLower(args[3])
	idsBytes, err := APIstub.GetState(contract_id)
	if err != nil {
		return shim.Error(err.Error())
	} else if idsBytes != nil {
		fmt.Println("This contract_id already exists: " + contract_id)
		return shim.Error("This contract_id already exists: " + contract_id)
	}
	
	var trade = Trade{
		Os:Osform{From:exporter, To:importer, Doc_hash:os_hash},
		Lcr:Lcrform{From:"", To:"", Doc_hash:""},
		Lc:Lcform{From:"", To:"", Doc_hash:""},
		Sr:Srform{From:"", To:"", Doc_hash:""},
		Ci:Ciform{From:"", To:"", Doc_hash:""},
		Bl:Blform{From:"", To:"", Doc_hash:""},
		Do:Doform{From:"", To:"", Doc_hash:""}}
	
	tradeJSONasBytes, err := json.Marshal(trade)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = APIstub.PutState(contract_id,tradeJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}


func (s *SmartContract) recordLCR(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//  0-contract_id  1-importer  2-bank  3-lcr_hash
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. you should input contract_id, From, To, lcr_hash")
	}

	contract_id := args[0]
	importer := strings.ToLower(args[1])
	bank := strings.ToLower(args[2])
	lcr_hash := strings.ToLower(args[3])
	fmt.Println("- LCR record ", contract_id)

	tradeAsBytes, err := APIstub.GetState(contract_id)
	if err != nil {
		return shim.Error("Failed to Record:" + err.Error())
	} else if tradeAsBytes == nil {
		return shim.Error("Could not found contract_id")
	}
	tradeRecord := Trade{}
	err = json.Unmarshal(tradeAsBytes, &tradeRecord)
	if err != nil{
		return shim.Error(err.Error())
	}

	if len(tradeRecord.Lcr.Doc_hash) == 0 && len(tradeRecord.Os.Doc_hash) != 0 {
		tradeRecord.Os.From = ""
		tradeRecord.Os.To = ""
		tradeRecord.Os.Doc_hash = ""
		tradeRecord.Lcr.From = importer
		tradeRecord.Lcr.To = bank
		tradeRecord.Lcr.Doc_hash = lcr_hash
		tradeRecord.Lc.From = ""
		tradeRecord.Lc.To = ""
		tradeRecord.Lc.Doc_hash = ""
		tradeRecord.Sr.From = ""
		tradeRecord.Sr.To = ""
		tradeRecord.Sr.Doc_hash = ""
		tradeRecord.Ci.From = ""
		tradeRecord.Ci.To = ""
		tradeRecord.Ci.Doc_hash = ""
		tradeRecord.Bl.From = ""
		tradeRecord.Bl.To = ""
		tradeRecord.Bl.Doc_hash = ""
		tradeRecord.Do.From = ""
		tradeRecord.Do.To = ""
		tradeRecord.Do.Doc_hash = ""
		tradeAsBytes, _ = json.Marshal(tradeRecord)
		err = APIstub.PutState(contract_id, tradeAsBytes)
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to input LCR: %s", args[0]))
		}
		fmt.Println("LCR Recorded (success)")
		return shim.Success(nil)
	} else {
		return shim.Error("Already exists")
	}
}	


func (s *SmartContract) recordLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//  0-contract_id  1-bank  2-exporter  3-lc_hash
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. you should input contract_id, From, To, lc_hash")
	}

	contract_id := args[0]
	bank := strings.ToLower(args[1])
	exporter := strings.ToLower(args[2])
	lc_hash := strings.ToLower(args[3])
	fmt.Println("- LC record ", contract_id)

	tradeAsBytes, err := APIstub.GetState(contract_id)
	if err != nil {
		return shim.Error("Failed to Record:" + err.Error())
	} else if tradeAsBytes == nil {
		return shim.Error("Could not found contract_id")
	}
	tradeRecord := Trade{}
	err = json.Unmarshal(tradeAsBytes, &tradeRecord)
	if err != nil{
		return shim.Error(err.Error())
	}

	if len(tradeRecord.Lc.Doc_hash) == 0 && len(tradeRecord.Lcr.Doc_hash) != 0 {
		tradeRecord.Os.From = ""
		tradeRecord.Os.To = ""
		tradeRecord.Os.Doc_hash = ""
		tradeRecord.Lcr.From = ""
		tradeRecord.Lcr.To = ""
		tradeRecord.Lcr.Doc_hash = ""
		tradeRecord.Lc.From = bank
		tradeRecord.Lc.To = exporter
		tradeRecord.Lc.Doc_hash = lc_hash
		tradeRecord.Sr.From = ""
		tradeRecord.Sr.To = ""
		tradeRecord.Sr.Doc_hash = ""
		tradeRecord.Ci.From = ""
		tradeRecord.Ci.To = ""
		tradeRecord.Ci.Doc_hash = ""
		tradeRecord.Bl.From = ""
		tradeRecord.Bl.To = ""
		tradeRecord.Bl.Doc_hash = ""
		tradeRecord.Do.From = ""
		tradeRecord.Do.To = ""
		tradeRecord.Do.Doc_hash = ""
		
		

		tradeAsBytes, _ = json.Marshal(tradeRecord)
		err = APIstub.PutState(contract_id, tradeAsBytes)
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to input LC: %s", args[0]))
		}
		fmt.Println("LC Recorded (success)")
		return shim.Success(nil)
	}else{
		return shim.Error("Already exists")
	}
}	

func (s *SmartContract) recordSR(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//  0-contract_id  1-exporter  2-shipper  3-sr_hash
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. you should input contract_id, From, To, sr_hash")
	}

	contract_id := args[0]
	exporter := strings.ToLower(args[1])
	shipper := strings.ToLower(args[2])
	sr_hash := strings.ToLower(args[3])
	fmt.Println("- SR record ", contract_id)

	tradeAsBytes, err := APIstub.GetState(contract_id)
	if err != nil {
		return shim.Error("Failed to Record:" + err.Error())
	} else if tradeAsBytes == nil {
		return shim.Error("Could not found contract_id")
	}
	tradeRecord := Trade{}
	err = json.Unmarshal(tradeAsBytes, &tradeRecord)
	if err != nil{
		return shim.Error(err.Error())
	}
	
	if len(tradeRecord.Sr.Doc_hash) == 0 && len(tradeRecord.Lc.Doc_hash) != 0 {
		tradeRecord.Os.From = ""
		tradeRecord.Os.To = ""
		tradeRecord.Os.Doc_hash = ""
		tradeRecord.Lcr.From = ""
		tradeRecord.Lcr.To = ""
		tradeRecord.Lcr.Doc_hash = ""
		tradeRecord.Lc.From = ""
		tradeRecord.Lc.To = ""
		tradeRecord.Lc.Doc_hash = ""
		tradeRecord.Sr.From = exporter
		tradeRecord.Sr.To = shipper
		tradeRecord.Sr.Doc_hash = sr_hash
		tradeRecord.Ci.From = ""
		tradeRecord.Ci.To = ""
		tradeRecord.Ci.Doc_hash = ""
		tradeRecord.Bl.From = ""
		tradeRecord.Bl.To = ""
		tradeRecord.Bl.Doc_hash = ""
		tradeRecord.Do.From = ""
		tradeRecord.Do.To = ""
		tradeRecord.Do.Doc_hash = ""

		tradeAsBytes, _ = json.Marshal(tradeRecord)
		err = APIstub.PutState(contract_id, tradeAsBytes)
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to input SR: %s", args[0]))
		}
		fmt.Println("SR Recorded (success)")
		return shim.Success(nil)
	}else{
		return shim.Error("Already exists")
	}
}	


func (s *SmartContract) recordBL(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//  0-contract_id  1-shipper  2-importer  3-bl_hash
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. you should input contract_id, From, To, bl_hash")
	}

	contract_id := args[0]
	shipper := strings.ToLower(args[1])
	importer := strings.ToLower(args[2])
	bl_hash := strings.ToLower(args[3])
	fmt.Println("- BL record ", contract_id)

	tradeAsBytes, err := APIstub.GetState(contract_id)
	if err != nil {
		return shim.Error("Failed to Record:" + err.Error())
	} else if tradeAsBytes == nil {
		return shim.Error("Could not found contract_id")
	}
	tradeRecord := Trade{}
	err = json.Unmarshal(tradeAsBytes, &tradeRecord)
	if err != nil{
		return shim.Error(err.Error())
	}

	if len(tradeRecord.Bl.Doc_hash) == 0 && len(tradeRecord.Sr.Doc_hash) != 0 {
		tradeRecord.Os.From = ""
		tradeRecord.Os.To = ""
		tradeRecord.Os.Doc_hash = ""
		tradeRecord.Lcr.From = ""
		tradeRecord.Lcr.To = ""
		tradeRecord.Lcr.Doc_hash = ""
		tradeRecord.Lc.From = ""
		tradeRecord.Lc.To = ""
		tradeRecord.Lc.Doc_hash = ""
		tradeRecord.Sr.From = ""
		tradeRecord.Sr.To = ""
		tradeRecord.Sr.Doc_hash = ""
		tradeRecord.Ci.From = ""
		tradeRecord.Ci.To = ""
		tradeRecord.Ci.Doc_hash = ""
		tradeRecord.Bl.From = shipper
		tradeRecord.Bl.To = importer
		tradeRecord.Bl.Doc_hash = bl_hash
		tradeRecord.Do.From = ""
		tradeRecord.Do.To = ""
		tradeRecord.Do.Doc_hash = ""

		tradeAsBytes, _ = json.Marshal(tradeRecord)
		err = APIstub.PutState(contract_id, tradeAsBytes)
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to input BL: %s", args[0]))
		}
		fmt.Println("BL Recorded (success)")
		return shim.Success(nil)
	}else{
		return shim.Error("Already exists")
	}
}	

func (s *SmartContract) recordCI(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//  0-contract_id  1-exporter  2-importer  3-ci_hash
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. you should input contract_id, From, To, ci_hash")
	}

	contract_id := args[0]
	exporter := strings.ToLower(args[1])
	importer := strings.ToLower(args[2])
	ci_hash := strings.ToLower(args[3])
	fmt.Println("- CI record ", contract_id)

	tradeAsBytes, err := APIstub.GetState(contract_id)
	if err != nil {
		return shim.Error("Failed to Record:" + err.Error())
	} else if tradeAsBytes == nil {
		return shim.Error("Could not found contract_id")
	}
	tradeRecord := Trade{}
	err = json.Unmarshal(tradeAsBytes, &tradeRecord)
	if err != nil{
		return shim.Error(err.Error())
	}
	
	if len(tradeRecord.Ci.Doc_hash) == 0 && len(tradeRecord.Bl.Doc_hash) != 0 {
		tradeRecord.Os.From = ""
		tradeRecord.Os.To = ""
		tradeRecord.Os.Doc_hash = ""
		tradeRecord.Lcr.From = ""
		tradeRecord.Lcr.To = ""
		tradeRecord.Lcr.Doc_hash = ""
		tradeRecord.Lc.From = ""
		tradeRecord.Lc.To = ""
		tradeRecord.Lc.Doc_hash = ""
		tradeRecord.Sr.From = ""
		tradeRecord.Sr.To = ""
		tradeRecord.Sr.Doc_hash = ""
		tradeRecord.Ci.From = exporter
		tradeRecord.Ci.To = importer
		tradeRecord.Ci.Doc_hash = ci_hash
		tradeRecord.Bl.From = ""
		tradeRecord.Bl.To = ""
		tradeRecord.Bl.Doc_hash = ""
		tradeRecord.Do.From = ""
		tradeRecord.Do.To = ""
		tradeRecord.Do.Doc_hash = ""

		tradeAsBytes, _ = json.Marshal(tradeRecord)
		err = APIstub.PutState(contract_id, tradeAsBytes)
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to input CI: %s", args[0]))
		}
		fmt.Println("CI Recorded (success)")
		return shim.Success(nil)
	}else{
		return shim.Error("Already exists")
	}
}

func (s *SmartContract) recordDO(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	//  0-contract_id  1-shipper  2-importer  3-do_hash
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. you should input contract_id, From, To, do_hash")
	}

	contract_id := args[0]
	shipper := strings.ToLower(args[1])
	importer := strings.ToLower(args[2])
	do_hash := strings.ToLower(args[3])
	fmt.Println("- DO record ", contract_id)

	tradeAsBytes, err := APIstub.GetState(contract_id)
	if err != nil {
		return shim.Error("Failed to Record:" + err.Error())
	} else if tradeAsBytes == nil {
		return shim.Error("Could not found contract_id")
	}
	tradeRecord := Trade{}
	err = json.Unmarshal(tradeAsBytes, &tradeRecord)
	if err != nil{
		return shim.Error(err.Error())
	}
	
	if len(tradeRecord.Do.Doc_hash) == 0 && len(tradeRecord.Ci.Doc_hash) != 0 {
		tradeRecord.Os.From = ""
		tradeRecord.Os.To = ""
		tradeRecord.Os.Doc_hash = ""
		tradeRecord.Lcr.From = ""
		tradeRecord.Lcr.To = ""
		tradeRecord.Lcr.Doc_hash = ""
		tradeRecord.Lc.From = ""
		tradeRecord.Lc.To = ""
		tradeRecord.Lc.Doc_hash = ""
		tradeRecord.Sr.From = ""
		tradeRecord.Sr.To = ""
		tradeRecord.Ci.From = ""
		tradeRecord.Ci.To = ""
		tradeRecord.Ci.Doc_hash = ""
		tradeRecord.Sr.Doc_hash = ""
		tradeRecord.Bl.From = ""
		tradeRecord.Bl.To = ""
		tradeRecord.Bl.Doc_hash = ""
		tradeRecord.Do.From = shipper
		tradeRecord.Do.To = importer
		tradeRecord.Do.Doc_hash = do_hash

		tradeAsBytes, _ = json.Marshal(tradeRecord)
		err = APIstub.PutState(contract_id, tradeAsBytes)
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to input DO: %s", args[0]))
		}
		fmt.Println("DO Recorded (success)")
		return shim.Success(nil)
	}else{
		return shim.Error("Already exists")
	}
}	

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
