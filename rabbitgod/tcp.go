package rabbitgod

import (
	//"io"
	"net"

	"github.com/xww/rabbitgo/internal/protocol"
)

type tcpServer struct {
	ctx *context
}

func (p *tcpServer) Handle(clientConn net.Conn) {
	p.ctx.rabbitgod.logf("TCP: new client(%s)", clientConn.RemoteAddr())


	var prot protocol.Protocol
	prot = &protocolV2{ctx: p.ctx}


	err := prot.IOLoop(clientConn)
	if err != nil {
		p.ctx.rabbitgod.logf("ERROR: client(%s) - %s", clientConn.RemoteAddr(), err)
		return
	}
}
