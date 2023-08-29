package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircledEquals(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="26.68" fill="#fff" fill-rule="evenodd"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><circle cx="36" cy="36" r="26.68" stroke-width="4.74"/><path stroke-width="8.031" d="M28.03 42.18h15.95M28.03 29.82h15.95" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}