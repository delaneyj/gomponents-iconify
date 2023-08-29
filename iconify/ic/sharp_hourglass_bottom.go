package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpHourglassBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 22l-.01-6L14 12l3.99-4.01L18 2H6v6l4 4l-4 3.99V22h12zM8 7.5V4h8v3.5l-4 4l-4-4z"/>`),
		g.Group(children),
	)
}