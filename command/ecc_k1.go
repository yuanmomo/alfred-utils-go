package command

// Package is called aw
import (
	"crypto/ecdsa"
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

type ECCk1Command struct{}

func (c *ECCk1Command) Name() string {
	return "k1"
}

var k1SubCommand []string = []string{"new", "pb"}

func (c *ECCk1Command) Description() Description {
	return Description{
		Short: "ECC key operation of k1 curve.",
		Usage: []string{
			"k1 new",
			"k1 pb string",
		},
	}
}

func (c *ECCk1Command) Execute(wf *aw.Workflow, args []string) *aw.Workflow {
	if len(args) < 1 {
		return wf
	}

	// encode or decode
	optionType := "new"
	// input string
	inputString := ""

	optionType = args[0]
	if len(args) == 2 {
		inputString = args[1]
	}

	var privateKey *ecdsa.PrivateKey
	var err error

	switch strings.TrimSpace(optionType) {
	case "new":
		privateKey, err = crypto.GenerateKey()
		if err != nil {
			result := "K1 key pair generate failed!!!"
			wf.NewItem(result).Valid(true).Arg(result).Subtitle(inputString)
		}
		addItem(wf, privateKey)
	case "pb":
		privateKey, err = crypto.HexToECDSA(inputString)
		if err != nil {
			result := "Get k1 private key from hex string failed!!!"
			wf.NewItem(result).Valid(true).Arg(result).Subtitle(inputString)
		}
		addItem(wf, privateKey)
	}

	return wf
}

// append private key and public key
func addItem(wf *aw.Workflow, privateKey *ecdsa.PrivateKey) *aw.Workflow {
	hexPrivateKey := hexutil.EncodeBig(privateKey.D)
	hexPrivateKey = remove0x(hexPrivateKey)

	// compress public key
	hexPublicKeyCompress := hexutil.Encode(crypto.CompressPubkey(&privateKey.PublicKey))
	hexPublicKeyCompress = remove0x(hexPublicKeyCompress)
	wf.NewItem(hexPublicKeyCompress).Valid(true).Arg(fmt.Sprintf("%s  %s", hexPrivateKey, hexPublicKeyCompress)).Subtitle(hexPrivateKey)

	// compress public key
	hexPublicKeyDecompress := fmt.Sprintf("04%s%s", remove0x(hexutil.EncodeBig(privateKey.PublicKey.X)),
		remove0x(hexutil.EncodeBig(privateKey.PublicKey.Y)))
	wf.NewItem(hexPublicKeyDecompress).Valid(true).Arg(fmt.Sprintf("%s  %s", hexPrivateKey, hexPublicKeyDecompress)).Subtitle(hexPrivateKey)
	return wf
}

func remove0x(input string) string{
	return strings.TrimPrefix(input,"0x")
}

func init() {
	RegisterCommand(&ECCk1Command{})
}
