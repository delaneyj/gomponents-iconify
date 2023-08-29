package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DimButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="5" fill="#fcea2b"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="5"/><path d="M36 26.087V24m7.01 4.99l1.475-1.475M45.913 36H48m-4.99 7.01l1.475 1.475M36 45.913V48m-7.01-4.99l-1.475 1.475M26.087 36H24m4.99-7.01l-1.475-1.475"/></g>`),
		g.Group(children),
	)
}