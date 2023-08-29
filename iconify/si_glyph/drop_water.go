package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropWater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 10.333C3 13.463 5.427 16 8.418 16C11.41 16 14 13.463 14 10.333C14 7.204 8.418 0 8.418 0S3 7.204 3 10.333z"/>`),
		g.Group(children),
	)
}