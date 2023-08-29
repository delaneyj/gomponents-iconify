package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewInAr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 1.5h5v2h-3v3h-2v-5Zm16 0h5v5h-2v-3h-3v-2ZM12 3.845l7.061 4.077v8.155L12 20.155l-7.063-4.078l.002-8.155l7.06-4.077Zm-5.062 6.386v4.692L11 17.268v-4.691L6.938 10.23ZM13 17.268l4.062-2.345v-4.691L13 12.577v4.69Zm-1-6.423L16.061 8.5L12 6.155L7.938 8.499L12 10.845ZM3.5 17.5v3h3v2h-5v-5h2Zm19 0v5h-5v-2h3v-3h2Z"/>`),
		g.Group(children),
	)
}