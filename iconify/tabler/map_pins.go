package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPins(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.828 9.828a4 4 0 1 0-5.656 0L8 12.657l2.828-2.829zM8 7v.01m10.828 10.818a4 4 0 1 0-5.656 0L16 20.657l2.828-2.829zM16 15v.01"/>`),
		g.Group(children),
	)
}