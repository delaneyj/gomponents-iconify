package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alignadjust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M29 50h609c15 0 28 13 28 27v41c0 16-13 29-28 29H29c-16 0-29-13-29-29V77c0-15 13-27 29-27zm0 166h609c15 0 28 12 28 26v42c0 15-13 28-28 28H29c-16 0-29-13-29-28v-42c0-14 13-26 29-26zm0 166h609c15 0 28 12 28 26v42c0 15-13 28-28 28H29c-16 0-29-13-29-28v-42c0-14 13-26 29-26zm0 164h609c15 0 28 14 28 28v41c0 16-13 29-28 29H29c-16 0-29-13-29-29v-41c0-15 13-28 29-28z"/>`),
		g.Group(children),
	)
}