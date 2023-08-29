package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TradeMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="6.735" d="M57.68 47.79V24.22l-10.1 20.2l-10.1-20.2v23.57M14.32 24.21h13.47m-6.74 0v23.57" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}