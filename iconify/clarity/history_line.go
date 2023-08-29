package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 9.83a1 1 0 0 0-1 1v8.72l5.9 4a1 1 0 0 0 1.1-1.67l-5-3.39v-7.66a1 1 0 0 0-1-1Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M18 2a16.09 16.09 0 0 0-14 8.26V5.2a1 1 0 0 0-2 0V14h8.8a1 1 0 0 0 0-2H5.35a14 14 0 1 1 3.23 16.35a1 1 0 0 0-1.35 1.48A16 16 0 1 0 18 2Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}