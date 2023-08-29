package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpDownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="m51.765 50.263l-4.035-3.71l-8.668 9.208V16.238l8.668 9.207l4.035-3.709L36.002 5L20.235 21.736l4.033 3.709l8.665-9.201v39.511l-8.665-9.202l-4.033 3.71l15.767 16.736z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m51.765 50.263l-4.035-3.71l-8.668 9.208V16.238l8.668 9.207l4.035-3.709L36.002 5L20.235 21.736l4.033 3.709l8.665-9.201v39.511l-8.665-9.202l-4.033 3.71l15.767 16.736z"/>`),
		g.Group(children),
	)
}