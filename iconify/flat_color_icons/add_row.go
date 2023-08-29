package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddRow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="M43 30V18c0-2.2-1.8-4-4-4H9c-2.2 0-4 1.8-4 4v12c0 2.2 1.8 4 4 4h30c2.2 0 4-1.8 4-4zM9 18h30v12H9V18z"/><circle cx="38" cy="38" r="10" fill="#43A047"/><g fill="#fff"><path d="M32 36h12v4H32z"/><path d="M36 32h4v12h-4z"/></g>`),
		g.Group(children),
	)
}