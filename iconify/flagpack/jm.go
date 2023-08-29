package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackJm0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackJm1)"><use href="#flagpackJm0"/><path fill="#093" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#272727" stroke="#FECA00" stroke-width="2.7" d="m-.14-1.041l-2.21-1.824v29.73l2.21-1.824l14.537-12L15.658 12l-1.26-1.041l-14.539-12Z"/><path fill="#272727" stroke="#FECA00" stroke-width="2.7" d="m32.164-1.06l2.186-1.724v29.568l-2.186-1.724l-15.219-12L15.601 12l1.344-1.06l15.22-12Z"/></g><defs><clipPath id="flagpackJm1"><use href="#flagpackJm0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}