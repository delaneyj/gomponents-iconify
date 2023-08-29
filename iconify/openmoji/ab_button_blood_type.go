package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbButtonBloodType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiAbButtonBloodType0" d="M59.035 60.137h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><path id="openmojiAbButtonBloodType1" d="m33.287 43.637l-6.429-15l-6.428 15m2.142-3.591h8.572m15.511-3.909h-5.72v-7.472h5.72a3.736 3.736 0 0 1 3.736 3.736a3.736 3.736 0 0 1-3.736 3.736Zm0 7.473h-5.72v-7.473h5.72a3.736 3.736 0 0 1 3.736 3.736a3.736 3.736 0 0 1-3.736 3.736Z"/></defs><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><use href="#openmojiAbButtonBloodType0"/><use href="#openmojiAbButtonBloodType1" stroke-linecap="round"/></g><path fill="#d22f27" d="M59.035 60.137h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="2"><use href="#openmojiAbButtonBloodType0"/><use href="#openmojiAbButtonBloodType1" stroke-linecap="round"/></g>`),
		g.Group(children),
	)
}