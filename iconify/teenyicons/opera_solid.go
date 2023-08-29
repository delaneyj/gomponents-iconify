package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OperaSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 7.5a7.5 7.5 0 0 1 13.146-4.936A17.5 17.5 0 0 0 8.741 2H7.5A3.5 3.5 0 0 0 4 5.5v4A3.5 3.5 0 0 0 7.5 13h1.241c1.488 0 2.969-.19 4.405-.563A7.5 7.5 0 0 1 0 7.5Z"/><path fill="currentColor" d="M14.073 11.115A7.466 7.466 0 0 0 15 7.5c0-1.31-.336-2.543-.927-3.615l-.114-.038a16.5 16.5 0 0 0-3.962-.8A3.489 3.489 0 0 1 11 5.5v4a3.49 3.49 0 0 1-1.003 2.452a16.499 16.499 0 0 0 3.962-.799l.114-.038Z"/>`),
		g.Group(children),
	)
}