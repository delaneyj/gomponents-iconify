package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 208a8 8 0 0 1-8 8H32a8 8 0 1 1 0-16h192a8 8 0 0 1 8 8ZM40.09 168h175.82a16.1 16.1 0 0 0 12.48-26.23L146.74 40.94a24.11 24.11 0 0 0-37.48 0L27.61 141.77A16.1 16.1 0 0 0 40.09 168Z"/>`),
		g.Group(children),
	)
}