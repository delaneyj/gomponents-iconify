package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07c0-.53.434-.965.965-.965h46.07c.53 0 .965.434.965.965v46.07c0 .53-.434.965-.965.965z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><path d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07c0-.53.434-.965.965-.965h46.07c.53 0 .965.434.965.965v46.07c0 .53-.434.965-.965.965z"/><path stroke-linecap="round" stroke-linejoin="round" d="M37.486 32.47h-4.572v8h4.572m-4.572-4h3.429M23 40.47v-8l5.714 8v-8m21.286 0l-2.286 8l-2.285-8l-2.286 8l-2.286-8"/></g>`),
		g.Group(children),
	)
}