package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OkButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07c0-.53.434-.965.965-.965h46.07c.53 0 .965.434.965.965v46.07c0 .53-.434.965-.965.965z"/><g fill="none" stroke="#000" stroke-miterlimit="10"><path stroke-width="2" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07c0-.53.434-.965.965-.965h46.07c.53 0 .965.434.965.965v46.07c0 .53-.434.965-.965.965z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M39.32 29.85v12.3m0-4.32l7.03-7.98m-3.64 4.14l4.21 8.16m-17.45 0v0a4.393 4.393 0 0 1-4.393-4.393v-3.515a4.393 4.393 0 0 1 4.393-4.393v0a4.393 4.393 0 0 1 4.393 4.393v3.515a4.393 4.393 0 0 1-4.393 4.393z"/></g>`),
		g.Group(children),
	)
}