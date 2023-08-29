package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBb0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBb1)"><use href="#flagpackBb0"/><path fill="#2E42A5" fill-rule="evenodd" d="M22 0h10v24H22V0Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M10 0h12v24H10V0Z" clip-rule="evenodd"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0h10v24H0V0Z" clip-rule="evenodd"/><path fill="#000" fill-rule="evenodd" d="M18.612 15.668c.188-1.29 1.52-4.093 1.52-4.093l.619-1.644l-2.45.795l.58.301l-1.135 3.199l-.856-.205V7.853l.677-.05l-1.672-2.431l-1.632 2.48h.787v6.17l-.965.16l-.688-3.234l.397-.333l-2.665-.492l.796 1.469s1.07 2.516 1.472 4.076l1.653-.128v2.398h1.84V15.54l1.722.128Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackBb1"><use href="#flagpackBb0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}