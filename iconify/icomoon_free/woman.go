package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Woman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 1.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 9 1.5zM11.234 8L12 7.445L9.917 4.224a.5.5 0 0 0-.417-.225h-4a.497.497 0 0 0-.417.225L3 7.445L3.766 8l1.729-2.244l.601 1.402l-2.095 3.841h1.917l.333 5h1v-5h.5v5h1l.333-5h1.917L8.906 7.157l.601-1.402l1.729 2.244z"/>`),
		g.Group(children),
	)
}