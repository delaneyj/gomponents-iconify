package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plugin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M10 7.25C10 5.75 9.532 5 7.971 5C6.471 5 6 5.812 6 7.25V8H5c-1.56 0-2 .75-2 2.25v6c0 1.529.44 1.792 2 1.792h14c1.563 0 2.004-.263 2.004-1.792v-6C21.004 8.75 20.563 8 19 8h-1v-.75C18 5.812 17.471 5 15.971 5C14.409 5 14 5.75 14 7.25V8h-4v-.75zm0 0"/>`),
		g.Group(children),
	)
}