package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 52a32 32 0 1 0 32 32a32 32 0 0 0-32-32Zm0 40a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm-52-36a52 52 0 0 0-104 0v94.69a64 64 0 1 0 104 0Zm-52 172a40 40 0 0 1-30.91-65.39a12 12 0 0 0 2.91-7.83V56a28 28 0 0 1 56 0v98.77a12 12 0 0 0 2.77 7.68A40 40 0 0 1 108 228Zm24-40a24 24 0 1 1-36-20.78V92a12 12 0 0 1 24 0v75.22A24 24 0 0 1 132 188Z"/>`),
		g.Group(children),
	)
}