package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldCheckeredFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M11.013 12v9.754A13 13 0 0 1 2.28 12h8.734zm9.284 3.794a13 13 0 0 1-7.283 5.951L13.013 12h8.708a12.96 12.96 0 0 1-1.424 3.794zM11.014 2.526L11.013 10H2.027c-.068-1.432.101-2.88.514-4.282a1 1 0 0 1 1.005-.717a11 11 0 0 0 7.192-2.256l.276-.219zM13.013 10V2.547l-.09-.073a11 11 0 0 0 7.189 2.537l.342-.01a1 1 0 0 1 1.005.717c.413 1.403.582 2.85.514 4.282h-8.96z"/></g>`),
		g.Group(children),
	)
}