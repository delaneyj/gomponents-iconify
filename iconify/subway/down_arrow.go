package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M289.7 341.3V0h-85.4v341.3L33.7 170.7v128L247 512l213.3-213.3v-128z"/>`),
		g.Group(children),
	)
}