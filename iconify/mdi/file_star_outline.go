package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileStarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2H6c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V8l-6-6m4 18H6V4h7v5h5v11m-8.5-2l.7-2.8L8 13.3l2.9-.2l1.1-2.7l1.1 2.6l2.9.2l-2.2 1.9l.7 2.8l-2.5-1.4L9.5 18Z"/>`),
		g.Group(children),
	)
}