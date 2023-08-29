package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OButtonBloodType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiOButtonBloodType0" stroke-miterlimit="10" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><path id="openmojiOButtonBloodType1" stroke-linecap="round" stroke-linejoin="round" d="M35.75 46.512a7.857 7.857 0 0 1-7.857-7.857v-6.286a7.857 7.857 0 0 1 7.857-7.857a7.857 7.857 0 0 1 7.857 7.857v6.286a7.857 7.857 0 0 1-7.857 7.857Z"/></defs><g fill="none" stroke="#000" stroke-width="2"><use href="#openmojiOButtonBloodType0" stroke-miterlimit="10"/><use href="#openmojiOButtonBloodType1" stroke-linecap="round" stroke-linejoin="round"/></g><path fill="#d22f27" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><g fill="none" stroke="#fff" stroke-width="2"><use href="#openmojiOButtonBloodType0" stroke-miterlimit="10"/><use href="#openmojiOButtonBloodType1" stroke-linecap="round" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}