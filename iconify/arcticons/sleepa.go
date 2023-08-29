package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleepa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.18 34.32A21.493 21.493 0 0 1 15.454 4.275a21.498 21.498 0 1 0 28.271 28.27a21.413 21.413 0 0 1-8.545 1.775Z"/>`),
		g.Group(children),
	)
}