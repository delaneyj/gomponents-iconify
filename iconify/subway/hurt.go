package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hurt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 28.3c-64 0-96.2 27.6-128 64c-31.8-36.4-64-64-128-64S0 71 0 199c0 64 64 192 256 298.7C448 391 512 263 512 199c0-128-64-170.7-128-170.7z"/>`),
		g.Group(children),
	)
}