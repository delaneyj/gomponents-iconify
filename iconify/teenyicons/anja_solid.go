package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnjaSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8 13a2 2 0 0 0 2-2H9a1 1 0 0 1-1 1H6.5v1H8Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.5 15a5.14 5.14 0 0 1-.543-.029L7 15H1V6h.019A6.5 6.5 0 0 1 14 6.5V15H7.5ZM8.974 4.9l1.3.086A5.951 5.951 0 0 1 11.503 7H4.207a5.616 5.616 0 0 1 4.767-2.1ZM9.085 8h2.71l.04.196a4 4 0 0 1 .07.804H10.5a1.5 1.5 0 0 1-1.415-1Zm1.415 2h1.286a8.016 8.016 0 0 1-.524 1.556a4.116 4.116 0 0 1-7.524 0A6.23 6.23 0 0 1 3.254 10H4.5a2.5 2.5 0 0 0 2.45-2h1.1a2.5 2.5 0 0 0 2.45 2Zm-6-1H3.235c.032-.23.09-.453.18-.669c.048-.112.099-.223.153-.331h2.347A1.5 1.5 0 0 1 4.5 9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}