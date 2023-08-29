package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.279 17.039a7.926 7.926 0 0 1-5.206 1.941a7.964 7.964 0 0 1-7.96-7.96a7.964 7.964 0 0 1 7.96-7.96a7.964 7.964 0 0 1 7.96 7.96a7.93 7.93 0 0 1-2.04 5.319l.165.165l3.583 3.582c.455.456-.252 1.163-.707.708l-3.755-3.755Zm1.754-6.019a6.964 6.964 0 0 0-6.96-6.96a6.963 6.963 0 0 0-6.96 6.96a6.963 6.963 0 0 0 6.96 6.96a6.964 6.964 0 0 0 6.96-6.96Zm-4.96-.5c.645 0 .643 1 0 1h-4c-.645 0-.643-1 0-1h4Z"/>`),
		g.Group(children),
	)
}