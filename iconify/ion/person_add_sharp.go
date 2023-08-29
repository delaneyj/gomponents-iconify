package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonAddSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M106 304v-54h54v-36h-54v-54H70v54H16v36h54v54h36z"/><circle cx="288" cy="144" r="112" fill="currentColor"/><path fill="currentColor" d="M288 288c-69.42 0-208 42.88-208 128v64h416v-64c0-85.12-138.58-128-208-128Z"/>`),
		g.Group(children),
	)
}