package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07c0-.53.434-.965.965-.965h46.07c.53 0 .965.434.965.965v46.07c0 .53-.434.965-.965.965z"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07c0-.53.434-.965.965-.965h46.07c.53 0 .965.434.965.965v46.07c0 .53-.434.965-.965.965z"/><path stroke-linecap="round" d="M30.11 47.005v-22h8.389a5.48 5.48 0 0 1 5.479 5.48h0a5.48 5.48 0 0 1-5.48 5.479H30.11"/></g>`),
		g.Group(children),
	)
}