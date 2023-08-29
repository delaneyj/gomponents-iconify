package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FortyTwoGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M320 96v320a160.019 160.019 0 0 0 113.138-46.862a160.019 160.019 0 0 0 34.683-174.367a160.006 160.006 0 0 0-86.591-86.592A160.019 160.019 0 0 0 320 96ZM0 256l160.002 160l160.001-160L160.002 96L0 256Zm480 0a160 160 0 0 0 160 160V96a160.002 160.002 0 0 0-160 160Z"/>`),
		g.Group(children),
	)
}