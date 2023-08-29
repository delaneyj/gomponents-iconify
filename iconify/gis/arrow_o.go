package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M14.291 2.346a5 5 0 0 0-2.477 4.944l7.556 81.087c.483 3.937 5.139 5.772 8.177 3.222l14.215-11.93l7.15 12.385c4.112 7.121 12.37 7.87 19.49 3.758c7.123-4.112 10.603-11.637 6.491-18.758l-7.125-12.341l17.164-6.245c3.727-1.357 4.466-6.305 1.298-8.692L19.784 2.688a5 5 0 0 0-5.493-.342Zm8.878 15.408l49.685 34.467l-19.713 7.173l.092.143l13 22.517c1.428 2.473-.356 3.67-2.83 5.098s-4.402 2.375-5.83-.098l-13-22.517l-.103-.198l-16.294 13.677z" color="currentColor"/>`),
		g.Group(children),
	)
}