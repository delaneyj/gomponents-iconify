package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TenThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiTenThirty0" d="M36 36.062v17m-.065-17.031l-9.526-5.5"/></defs><g fill="#FFF" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="35.958" cy="35.99" r="23"/><use href="#openmojiTenThirty0"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="35.958" cy="35.99" r="23"/><use href="#openmojiTenThirty0"/></g>`),
		g.Group(children),
	)
}