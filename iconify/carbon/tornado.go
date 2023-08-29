package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tornado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 26h4v2h-4zm-4-4h6v2h-6zm-4-4h10v2H8zm0-4h12v2H8zm2-4h14v2H10zM8 6h18v2H8z"/>`),
		g.Group(children),
	)
}