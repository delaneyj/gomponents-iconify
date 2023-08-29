package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" d="M52.498 49.249L36 67.036L19.502 49.249l4.076-3.789l9.641 10.395V5.036h5.562v50.819l9.641-10.395z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M52.498 49.249L36 67.036L19.502 49.249l4.076-3.789l9.641 10.395V5.036h5.562v50.819l9.641-10.395z"/>`),
		g.Group(children),
	)
}