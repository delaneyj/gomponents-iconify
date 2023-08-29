package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 11a1 1 0 0 0 .89-.55a1 1 0 0 0-.44-1.34l-2-1a1 1 0 1 0-.9 1.78l2 1A.93.93 0 0 0 10 11Zm2-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm-3.64-4.67a1 1 0 0 0-.13 1.4a1 1 0 0 0 1.41.13a3.76 3.76 0 0 1 4.72 0a1 1 0 0 0 .64.23a1 1 0 0 0 .64-1.76a5.81 5.81 0 0 0-7.28 0Zm7.19-7.22l-2 1a1 1 0 0 0-.44 1.34A1 1 0 0 0 14 11a.93.93 0 0 0 .45-.11l2-1a1 1 0 0 0-.9-1.78Z"/>`),
		g.Group(children),
	)
}