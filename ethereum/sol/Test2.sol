//
//  Panoramix v4 Oct 2019
//  Decompiled source of 0x9799b475dEc92Bd99bbdD943013325C36157f383
//
//  Let's make the world open source
//

def storage:
unknownbf353dbb is mapping of uint256 at storage 0
cacheAddress is addr at storage 1

def cache(): # not payable
return cacheAddress

def unknownbf353dbb(addr _param1): # not payable
require calldata.size - 4 >= 32
return unknownbf353dbb[_param1]

//
//Regular functions
//

def _fallback() payable: # default function
log 0xc1c40e9d: call.value, caller, this.address

def deny(address _operator): # not payable
require calldata.size - 4 >= 32
require unknownbf353dbb[caller] == 1
unknownbf353dbb[addr(_operator)] = 0

def unknown65fae35e(addr _param1): # not payable
require calldata.size - 4 >= 32
require unknownbf353dbb[caller] == 1
unknownbf353dbb[addr(_param1)] = 1

def setCache(address _cacheAddr): # not payable
require calldata.size - 4 >= 32
require unknownbf353dbb[caller] == 1
if not _cacheAddr:
revert with 0, 'ds-proxy-cache-address-required'
cacheAddress = _cacheAddr
return 1

def execute(address _target, bytes _data) payable:
require calldata.size - 4 >= 64
require _data <= 4294967296
require _data + 36 <= calldata.size
require _data.length <= 4294967296 and _data + _data.length + 36 <= calldata.size
require unknownbf353dbb[caller] == 1
delegate _target with:
gas gas_remaining wei
args _data[all]
require not delegate.return_code != 1

def unknown78e111f6(addr _param1, array _param2) payable:
require calldata.size - 4 >= 64
require _param2 <= 4294967296
require _param2 + 36 <= calldata.size
require _param2.length <= 4294967296 and _param2 + _param2.length + 36 <= calldata.size
require unknownbf353dbb[caller] == 1
delegate _param1 with:
gas gas_remaining wei
args _param2[all]
if not delegate.return_code == 1:
revert with ext_call.return_data[0 len return_data.size]
return Array(len=return_data.size, data=ext_call.return_data)

def execute(bytes _code, bytes _data) payable:
require calldata.size - 4 >= 64
require _code <= 4294967296
require _code + 36 <= calldata.size
require _code.length <= 4294967296 and _code + _code.length + 36 <= calldata.size
mem[128 len _code.length] = _code[all]
mem[_code.length + 128] = 0
require _data <= 4294967296
require _data + 36 <= calldata.size
require _data.length <= 4294967296 and _data + _data.length + 36 <= calldata.size
mem[ceil32(_code.length) + 128] = _data.length
mem[ceil32(_code.length) + 160 len _data.length] = _data[all]
mem[ceil32(_code.length) + _data.length + 160] = 0
mem[ceil32(_code.length) + ceil32(_data.length) + 164] = 32
mem[ceil32(_code.length) + ceil32(_data.length) + 196] = _code.length
mem[ceil32(_code.length) + ceil32(_data.length) + 228 len ceil32(_code.length)] = _code[all], mem[_code.length + 128 len ceil32(_code.length) - _code.length]
require ext_code.size(cacheAddress)
static call cacheAddress with:
gas gas_remaining wei
args Array(len=_code.length, data=_code[all])
mem[ceil32(_code.length) + ceil32(_data.length) + 160] = ext_call.return_data[0]
if not ext_call.success:
revert with ext_call.return_data[0 len return_data.size]
else:
require return_data.size >= 32
if addr(ext_call.return_data):
require unknownbf353dbb[caller] == 1
delegate ext_call.return_datamem[ceil32(_code.length) + 160 len 4] with:
gas gas_remaining wei
args mem[ceil32(_code.length) + 164 len Mask(8 * -ceil32(_code.length) + _code.length + 32, 0, 0), mem[_code.length + 160 len -_code.length + ceil32(_code.length)] - 4]
require not delegate.return_code != 1
stop
else:
mem[ceil32(_code.length) + ceil32(_data.length) + 164] = 32
mem[ceil32(_code.length) + ceil32(_data.length) + 196] = _code.length
mem[ceil32(_code.length) + ceil32(_data.length) + 228 len ceil32(_code.length)] = _code[all], mem[_code.length + 128 len ceil32(_code.length) - _code.length]
require ext_code.size(cacheAddress)
call cacheAddress with:
gas gas_remaining wei
args Array(len=_code.length, data=_code[all])
mem[ceil32(_code.length) + ceil32(_data.length) + 160] = ext_call.return_data[0]
if not ext_call.success:
revert with ext_call.return_data[0 len return_data.size]
else:
require return_data.size >= 32
require unknownbf353dbb[caller] == 1
delegate ext_call.return_datamem[ceil32(_code.length) + 160 len 4] with:
gas gas_remaining wei
args mem[ceil32(_code.length) + 164 len Mask(8 * -ceil32(_code.length) + _code.length + 32, 0, 0), mem[_code.length + 160 len -_code.length + ceil32(_code.length)] - 4]
require not delegate.return_code != 1
stop

def unknowna90e8731(array _param1, array _param2) payable:
require calldata.size - 4 >= 64
require _param1 <= 4294967296
require _param1 + 36 <= calldata.size
require _param1.length <= 4294967296 and _param1 + _param1.length + 36 <= calldata.size
mem[128 len _param1.length] = _param1[all]
mem[_param1.length + 128] = 0
require _param2 <= 4294967296
require _param2 + 36 <= calldata.size
require _param2.length <= 4294967296 and _param2 + _param2.length + 36 <= calldata.size
mem[ceil32(_param1.length) + 128] = _param2.length
mem[ceil32(_param1.length) + 160 len _param2.length] = _param2[all]
mem[ceil32(_param1.length) + _param2.length + 160] = 0
mem[ceil32(_param1.length) + ceil32(_param2.length) + 164] = 32
mem[ceil32(_param1.length) + ceil32(_param2.length) + 196] = _param1.length
mem[ceil32(_param1.length) + ceil32(_param2.length) + 228 len ceil32(_param1.length)] = _param1[all], mem[_param1.length + 128 len ceil32(_param1.length) - _param1.length]
require ext_code.size(cacheAddress)
static call cacheAddress with:
gas gas_remaining wei
args Array(len=_param1.length, data=_param1[all])
mem[ceil32(_param1.length) + ceil32(_param2.length) + 160] = ext_call.return_data[0]
if not ext_call.success:
revert with ext_call.return_data[0 len return_data.size]
else:
require return_data.size >= 32
if addr(ext_call.return_data):
require unknownbf353dbb[caller] == 1
delegate ext_call.return_datamem[ceil32(_param1.length) + 160 len 4] with:
gas gas_remaining wei
args mem[ceil32(_param1.length) + 164 len Mask(8 * -ceil32(_param1.length) + _param1.length + 32, 0, 0), mem[_param1.length + 160 len -_param1.length + ceil32(_param1.length)] - 4]
if not delegate.return_code == 1:
revert with ext_call.return_data[0 len return_data.size]
else:
return Array(len=return_data.size, data=ext_call.return_data)
else:
mem[ceil32(_param1.length) + ceil32(_param2.length) + 164] = 32
mem[ceil32(_param1.length) + ceil32(_param2.length) + 196] = _param1.length
mem[ceil32(_param1.length) + ceil32(_param2.length) + 228 len ceil32(_param1.length)] = _param1[all], mem[_param1.length + 128 len ceil32(_param1.length) - _param1.length]
require ext_code.size(cacheAddress)
call cacheAddress with:
gas gas_remaining wei
args Array(len=_param1.length, data=_param1[all])
mem[ceil32(_param1.length) + ceil32(_param2.length) + 160] = ext_call.return_data[0]
if not ext_call.success:
revert with ext_call.return_data[0 len return_data.size]
else:
require return_data.size >= 32
require unknownbf353dbb[caller] == 1
delegate ext_call.return_datamem[ceil32(_param1.length) + 160 len 4] with:
gas gas_remaining wei
args mem[ceil32(_param1.length) + 164 len Mask(8 * -ceil32(_param1.length) + _param1.length + 32, 0, 0), mem[_param1.length + 160 len -_param1.length + ceil32(_param1.length)] - 4]
if not delegate.return_code == 1:
revert with ext_call.return_data[0 len return_data.size]
else:
return Array(len=return_data.size, data=ext_call.return_data)

