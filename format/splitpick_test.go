package format

import (
	"testing"

	"github.com/trivago/gollum/core"
	"github.com/trivago/tgo/ttesting"
)

func TestSplitPick_Success(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.SplitPick")
	config.Override("Index", 0)
	config.Override("Delimiter", "#")
	plugin, err := core.NewPluginWithConfig(config)

	expect.NoError(err)

	formatter, casted := plugin.(*SplitPick)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte("MTIzNDU2#NjU0MzIx"), nil, core.InvalidStreamID)
	err = formatter.ApplyFormatter(msg)

	expect.NoError(err)
	expect.Equal("MTIzNDU2", msg.String())
}

func TestSplitPick_OutOfBoundIndex(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.SplitPick")
	config.Override("Index", 2)
	plugin, err := core.NewPluginWithConfig(config)

	expect.NoError(err)

	formatter, casted := plugin.(*SplitPick)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte("MTIzNDU2:NjU0MzIx"), nil, core.InvalidStreamID)
	err = formatter.ApplyFormatter(msg)

	expect.NoError(err)
	expect.Equal(0, len(msg.GetPayload()))
}

func TestSplitPickTarget(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.SplitPick")
	config.Override("Index", 0)
	config.Override("Delimiter", "#")
	config.Override("Target", "foo")
	plugin, err := core.NewPluginWithConfig(config)

	expect.NoError(err)

	formatter, casted := plugin.(*SplitPick)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte("PAYLOAD"), nil, core.InvalidStreamID)
	msg.GetMetadata().Set("foo", []byte("MTIzNDU2#NjU0MzIx"))
	err = formatter.ApplyFormatter(msg)

	expect.NoError(err)
	foo, err := msg.GetMetadata().Bytes("foo")
	expect.NoError(err)
	expect.Equal("PAYLOAD", msg.String())
	expect.Equal("MTIzNDU2", string(foo))
}
