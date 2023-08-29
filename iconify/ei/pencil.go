package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="m9.6 40.4l2.5-9.9L27 15.6l7.4 7.4l-14.9 14.9l-9.9 2.5zm4.3-8.9l-1.5 6.1l6.1-1.5L31.6 23L27 18.4L13.9 31.5z"/><path fill="currentColor" d="M17.8 37.3c-.6-2.5-2.6-4.5-5.1-5.1l.5-1.9c3.2.8 5.7 3.3 6.5 6.5l-1.9.5z"/><path fill="currentColor" d="m29.298 19.287l1.414 1.414l-13.01 13.02l-1.414-1.41zM11 39l2.9-.7c-.3-1.1-1.1-1.9-2.2-2.2L11 39zm24-16.6L27.6 15l3-3l.5.1c3.6.5 6.4 3.3 6.9 6.9l.1.5l-3.1 2.9zM30.4 15l4.6 4.6l.9-.9c-.5-2.3-2.3-4.1-4.6-4.6l-.9.9z"/>`),
		g.Group(children),
	)
}