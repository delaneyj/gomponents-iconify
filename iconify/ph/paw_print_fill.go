package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PawPrintFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 108a28 28 0 1 1-28-28a28 28 0 0 1 28 28Zm-168 0a28 28 0 1 0-28 28a28 28 0 0 0 28-28Zm20-20a28 28 0 1 0-28-28a28 28 0 0 0 28 28Zm72 0a28 28 0 1 0-28-28a28 28 0 0 0 28 28Zm23.12 60.86a35.3 35.3 0 0 1-16.87-21.14a44 44 0 0 0-84.5 0A35.25 35.25 0 0 1 69 148.82A40 40 0 0 0 88 224a39.48 39.48 0 0 0 15.52-3.13a64.09 64.09 0 0 1 48.87 0a40 40 0 0 0 34.73-72Z"/>`),
		g.Group(children),
	)
}