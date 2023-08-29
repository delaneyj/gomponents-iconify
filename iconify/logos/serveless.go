package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Serveless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#FFD426" d="M0 218.14l119.805-16.955l-105.676 173.49l240.742-217.569l-103.984 10.172L256 0z"/>`),
		g.Group(children),
	)
}