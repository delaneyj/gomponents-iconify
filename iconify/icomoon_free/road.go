package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 16h5L12 0H9l.5 4h-3L7 0H4L0 16h5l.5-4h5l.5 4zm-5.25-6l.5-4h3.5l.5 4h-4.5z"/>`),
		g.Group(children),
	)
}