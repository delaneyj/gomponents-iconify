package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aligncenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M500 147H167c-16 0-29-14-29-28V77c0-14 13-27 29-27h333c14 0 27 13 27 27v42c0 15-13 28-27 28zm110 162H55c-15 0-27-12-27-28v-41c0-16 12-27 27-27h555c15 0 28 11 28 27v41c0 16-13 28-28 28zm-83 163H139c-16 0-28-13-28-28v-42c0-15 12-28 28-28h388c14 0 28 13 28 28v42c0 15-14 28-28 28zM29 546h609c15 0 28 14 28 28v41c0 16-13 29-28 29H29c-16 0-29-13-29-29v-41c0-15 13-28 29-28z"/>`),
		g.Group(children),
	)
}