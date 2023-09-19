package singleByteXor

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/dmtaylor/cryptopals-cli/ciphers/xor"
	"github.com/dmtaylor/cryptopals-cli/util/distance"
	"github.com/dmtaylor/cryptopals-cli/util/distribution"
	"github.com/dmtaylor/cryptopals-cli/util/encoding"
)

const inputData = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
const targetDistribution = "etaoinshrdlu" // use abridged target dist

// SingleByteXorCmd represents the singleByteXor command
var SingleByteXorCmd = &cobra.Command{
	Use:   "singleByteXor",
	Short: "Solution command to Set 1 Challenge 3",
	RunE:  set1Challenge3Func,
}

type keyResult struct {
	key      byte
	distance int
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// SingleByteXorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// SingleByteXorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func set1Challenge3Func(_ *cobra.Command, _ []string) error {

	data, err := encoding.HexToBin(inputData)
	if err != nil {
		return err
	}
	distances := make([]keyResult, 0, 128)
	var i byte
	for i = 0; i < 128; i++ {
		result := xor.SingleByteXor(data, i)
		result = []byte(strings.ToLower(string(result)))
		dist := distribution.NewByteDistribution()
		for _, b := range result {
			_ = dist.Add(b, 1)
		}
		ord := dist.Ordering()
		d, err := distance.BitwiseHamming([]byte(targetDistribution), ord[0:len(targetDistribution)])
		if err != nil {
			return fmt.Errorf("err for dist on %d: %w", i, err)
		}
		res := keyResult{
			i,
			d,
		}
		distances = append(distances, res)
	}

	sort.SliceStable(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	for i = 0; i < 10; i++ {
		res := xor.SingleByteXor(data, distances[i].key)
		fmt.Printf("%q d(%d): %s\n", distances[i].key, distances[i].distance, string(res))
	}

	return nil
}
