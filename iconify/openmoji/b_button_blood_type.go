package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BButtonBloodType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiBButtonBloodType0" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><path id="openmojiBButtonBloodType1" stroke-linecap="round" d="M37.763 35.228h-8.42v-11h8.42a5.5 5.5 0 0 1 5.5 5.5a5.5 5.5 0 0 1-5.5 5.5Zm0 11h-8.42v-11h8.42a5.5 5.5 0 0 1 5.5 5.5a5.5 5.5 0 0 1-5.5 5.5Z"/></defs><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><use href="#openmojiBButtonBloodType0"/><use href="#openmojiBButtonBloodType1" stroke-linecap="round"/></g><path fill="#d22f27" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="2"><use href="#openmojiBButtonBloodType0"/><use href="#openmojiBButtonBloodType1" stroke-linecap="round"/></g>`),
		g.Group(children),
	)
}