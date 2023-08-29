package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCf0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCf1)"><use href="#flagpackCf0"/><path fill="#3D58DB" fill-rule="evenodd" d="M0 0h32v6H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 6h32v6H0V6Z" clip-rule="evenodd"/><path fill="#73BE4A" fill-rule="evenodd" d="M0 12h32v6H0v-6Z" clip-rule="evenodd"/><path fill="#FFD018" fill-rule="evenodd" d="M0 18h32v6H0v-6Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M4.53 5.416L2.104 7.098l.775-2.88l-1.78-1.84l2.41-.1L4.53-.57l1.019 2.848h2.406l-1.776 1.94l.89 2.71l-2.54-1.512Z" clip-rule="evenodd"/><path fill="#E11C1B" fill-rule="evenodd" d="M12 0h8v24h-8V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackCf1"><use href="#flagpackCf0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}