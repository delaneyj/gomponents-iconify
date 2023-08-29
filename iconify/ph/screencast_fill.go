package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreencastFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M56 208a8 8 0 0 1-16 0a8 8 0 0 0-8-8a8 8 0 0 1 0-16a24 24 0 0 1 24 24Zm-24-56a8 8 0 0 0 0 16a40 40 0 0 1 40 40a8 8 0 0 0 16 0a56.06 56.06 0 0 0-56-56Zm0-32a8 8 0 0 0 0 16a72.08 72.08 0 0 1 72 72a8 8 0 0 0 16 0a88.1 88.1 0 0 0-88-88Zm184-80H40a16 16 0 0 0-16 16v40a8 8 0 0 0 8 8a104.11 104.11 0 0 1 104 104a8 8 0 0 0 8 8h72a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}