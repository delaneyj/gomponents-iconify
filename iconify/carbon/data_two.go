package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 6h11v2H4zm14 0h10v2H18zm3 6h7v2h-7zm-10 0h7v2h-7zm-7 0h4v2H4zm0 6h24v2H4zm0 6h17v2H4zm20 0h4v2h-4z"/>`),
		g.Group(children),
	)
}