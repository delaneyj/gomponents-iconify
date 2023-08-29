package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Random(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 3v4h-4.594l-.281.5l-3.625 6.469L10.594 7H4v2h5.406l3.938 7l-3.938 7H4v2h6.594l9-16H23v4l5-5zm-6.219 15l-1.156 2.063L18.406 25H23v4l5-5l-5-5v4h-3.406z"/>`),
		g.Group(children),
	)
}