package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M20.5 4.5C9.454 4.5.5 11.887.5 21s8.954 16.5 20 16.5s20-7.387 20-16.5s-8.954-16.5-20-16.5zm-.031 30.791c-9.551 0-17.292-6.387-17.292-14.266c0-7.879 7.741-14.266 17.292-14.267c9.551 0 17.293 6.388 17.293 14.267c-.002 7.879-7.742 14.266-17.293 14.266z"/>`),
		g.Group(children),
	)
}