package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiClButton0" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><path id="openmojiClButton1" stroke-linecap="round" d="M32.07 42.578a5.314 5.314 0 0 1-3.538 1.343a5.334 5.334 0 0 1-5.334-5.334v-4.268a5.334 5.334 0 0 1 5.334-5.334a5.313 5.313 0 0 1 3.538 1.343m10.331-1.279v14.936h6.401"/></defs><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><use href="#openmojiClButton0"/><use href="#openmojiClButton1" stroke-linecap="round"/></g><path fill="#d22f27" d="M59.035 60.453h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="2"><use href="#openmojiClButton0"/><use href="#openmojiClButton1" stroke-linecap="round"/></g>`),
		g.Group(children),
	)
}