package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerStatAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M1.383 48.166a4.725 2.593 0 0 0 0 3.668L46.66 76.68a4.725 2.593 0 0 0 6.683 0l45.275-24.847a4.725 2.593 0 0 0 0-3.666L53.342 23.32a4.725 2.593 0 0 0-6.683-.001zM11.407 50l38.594-21.18L88.595 50L50 71.18Z" color="currentColor"/><path fill="currentColor" d="M42 35.674V58.59h6V35.674h-6zm20 5.59V58.59h6V41.264h-6zM32 42.51v16.08h6V42.51h-6zm20 4.66v11.42h6V47.17h-6zM32 60.59v4h36v-4H32z"/>`),
		g.Group(children),
	)
}