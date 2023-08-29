package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.325 3.05L8.667 20.432l1.932.518l4.658-17.382l-1.932-.518ZM7.612 18.36l1.36-1.448l-.001-.019l-5.094-4.78l4.79-5.105l-1.458-1.369l-6.16 6.563l6.563 6.159Zm8.776 0l-1.36-1.448l.001-.019l5.094-4.78l-4.79-5.105l1.458-1.369l6.16 6.563l-6.563 6.159Z"/>`),
		g.Group(children),
	)
}