package codec

import (
	"encoding/binary"

	"github.com/zjregee/anet"

	"github.com/cloudwego/netpoll-benchmark/runner"
)

func Encode(writer anet.Writer, msg *runner.Message) (err error) {
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, uint32(len(msg.Message)))
	writer.WriteBytes(header, 4)
	writer.WriteString(msg.Message, len(msg.Message))
	err = writer.Flush()
	return err
}

func Decode(reader anet.Reader, msg *runner.Message) (err error) {
	bLen, err := reader.ReadBytes(4)
	if err != nil {
		return err
	}
	l := int(binary.BigEndian.Uint32(bLen))
	msg.Message, err = reader.ReadString(l)
	if err != nil {
		return err
	}
	return err
}
