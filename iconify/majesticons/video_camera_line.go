package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 8a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v1.586l2.44-2.44c.944-.944 2.56-.275 2.56 1.061v7.586c0 1.336-1.616 2.006-2.56 1.06L17 14.415V16a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8zm15.414 4L20 14.586V9.414L17.414 12zM5 7a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H5z"/></g>`),
		g.Group(children),
	)
}