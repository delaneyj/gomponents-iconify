package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberThreeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M172 160a52 52 0 0 1-86.67 38.76a4 4 0 1 1 5.34-6A44 44 0 1 0 120 116a4 4 0 0 1-3.2-6.4L160 52H88a4 4 0 0 1 0-8h80a4 4 0 0 1 3.2 6.4l-43.61 58.15A52.08 52.08 0 0 1 172 160Z"/>`),
		g.Group(children),
	)
}