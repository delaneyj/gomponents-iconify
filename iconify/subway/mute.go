package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m159.8 320l64-64l-64-64l85.3-85.3l64 64V0L159.8 149.3H74.4v213.3h85.3L309.1 512V341.3l-64 64l-85.3-85.3zm245.3-128l-32-32l-64 64l-64-64l-32 32l64 64l-64 64l32 32l64-64l64 64l32-32l-64-64l64-64z"/>`),
		g.Group(children),
	)
}