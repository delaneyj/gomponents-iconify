package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bicycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiBicycle0" d="m21.1 46l1.7-21.4M36 45L23 30m22 0l8 16M23 30h21m2-1.6L36 45m0 0l17 1"/></defs><path d="M31 45a.945.945 0 0 0-1 1a9.2 9.2 0 1 1-9.2-9.2l.2-2h-.2A11.2 11.2 0 1 0 32 46a.945.945 0 0 0-1-1Zm21.8-10.2a12.271 12.271 0 0 0-3.9.7l.9 1.8a10.473 10.473 0 0 1 3-.5a9.2 9.2 0 1 1-9.1 9.7l-2-.1a11.155 11.155 0 1 0 11.1-11.6Z"/><path d="m48 38.2l-.9-1.8a11.093 11.093 0 0 0-5.3 7.9l2 .1a9.375 9.375 0 0 1 4.2-6.2Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23 22h4m15 3.7h6"/><use href="#openmojiBicycle0" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><use href="#openmojiBicycle0" fill="none" stroke="#ea5a47" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.1"/>`),
		g.Group(children),
	)
}