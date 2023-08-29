package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 56a28 28 0 1 0 28 28a28 28 0 0 0-28-28Zm0 40a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm-60 50.08V40a32 32 0 0 0-64 0v106.08a56 56 0 1 0 64 0ZM136 80h-32V40a16 16 0 0 1 32 0Z"/>`),
		g.Group(children),
	)
}