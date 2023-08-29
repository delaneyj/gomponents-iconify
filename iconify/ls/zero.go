package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M247 0c205 0 255 224 255 386s-50 386-255 386S0 546 0 386S41 0 247 0zm3 706c157 0 183-201 183-319c0-117-26-318-183-318C88 69 70 269 70 387s18 319 180 319z"/>`),
		g.Group(children),
	)
}