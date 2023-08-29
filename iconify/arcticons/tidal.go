package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tidal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.333 12.333l6.834 6.834L24 12.333l6.833 6.834l6.834-6.834l6.833 6.834L37.667 26l-6.834-6.833L24 26l6.833 6.833L24 39.667l-6.833-6.834L24 26l-6.833-6.833L10.333 26L3.5 19.167Z"/>`),
		g.Group(children),
	)
}