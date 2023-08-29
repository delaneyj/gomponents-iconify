package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftRightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="m50.263 20.234l-3.709 4.035l9.208 8.669H16.238l9.208-8.669l-3.709-4.035L5 35.998l16.737 15.767l3.709-4.034l-9.201-8.665h39.51l-9.201 8.665l3.709 4.034L67 35.998z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m50.263 20.234l-3.709 4.035l9.208 8.669H16.238l9.208-8.669l-3.709-4.035L5 35.998l16.737 15.767l3.709-4.034l-9.201-8.665h39.51l-9.201 8.665l3.709 4.034L67 35.998z"/>`),
		g.Group(children),
	)
}