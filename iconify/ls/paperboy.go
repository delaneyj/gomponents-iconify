package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperboy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M576 288C576 129 447 0 288 0S0 129 0 288c0 118 73 223 181 266c13 34 52 129 52 129c9 22 31 37 55 37s46-15 55-37c0 0 39-95 53-129c107-43 180-148 180-266zM288 601c-18-45-42-101-42-101c-2-7-7-12-14-13c-88-25-150-107-150-199c0-114 92-206 206-206s206 92 206 206c0 92-61 174-150 199c-6 1-12 6-14 13c0 0-23 56-42 101z"/>`),
		g.Group(children),
	)
}