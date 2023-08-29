package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lipstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M818.91 946.5q-47.5 47.5-94.5 67.5t-66 1l-457-458q-20-19 3-69l-166-166q-13-13-25.5-80T.41 128q0-28 16-60t32-50l16-18l299 330q47-20 67-1l457 457q19 19-1 66t-67.5 94.5z"/>`),
		g.Group(children),
	)
}