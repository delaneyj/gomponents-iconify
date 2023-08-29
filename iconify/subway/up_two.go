package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 63.6L0 319.6l69.8 69.8L256 203.2l186.2 186.2l69.8-69.8z"/>`),
		g.Group(children),
	)
}