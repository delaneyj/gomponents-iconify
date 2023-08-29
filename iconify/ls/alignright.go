package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alignright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M638 147H305c-15 0-27-14-27-28V77c0-14 12-27 27-27h333c15 0 28 13 28 27v42c0 15-13 28-28 28zm0 162H83c-15 0-28-12-28-28v-41c0-16 13-27 28-27h555c15 0 28 11 28 27v41c0 16-13 28-28 28zm0 163H250c-15 0-28-13-28-28v-42c0-15 13-28 28-28h388c15 0 28 13 28 28v42c0 15-13 28-28 28zM29 546h609c15 0 28 14 28 28v41c0 16-13 29-28 29H29c-16 0-29-13-29-29v-41c0-15 13-28 29-28z"/>`),
		g.Group(children),
	)
}