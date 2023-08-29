package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAudit0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="m8 36l.005-7.957a1 1 0 0 1 1-1h10.002c.922 0 .917-.818.917-2.764c0-1.946-4.902-3.585-4.902-10.426S20.1 5 24.32 5s8.817 2.012 8.817 8.853c0 6.841-4.876 7.929-4.876 10.426s0 2.764.78 2.764h9.96a1 1 0 0 1 1 1V36H8Z"/><path stroke-linecap="round" d="M8 42h32"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAudit0)"/>`),
		g.Group(children),
	)
}