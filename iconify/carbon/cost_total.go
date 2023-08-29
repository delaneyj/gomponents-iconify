package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CostTotal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 13h12v2H10zm0 5h12v2H10z"/><path fill="currentColor" d="M30 6.414L28.586 5l-3.35 3.35a12.081 12.081 0 0 0-1.586-1.586L27 3.414L25.586 2l-3.61 3.61A11.98 11.98 0 0 0 5.61 21.976L2 25.586L3.414 27l3.35-3.35a12.081 12.081 0 0 0 1.586 1.586L5 28.586L6.414 30l3.61-3.61A11.98 11.98 0 0 0 26.39 10.024ZM16 26a10 10 0 1 1 10-10a10.011 10.011 0 0 1-10 10Z"/>`),
		g.Group(children),
	)
}