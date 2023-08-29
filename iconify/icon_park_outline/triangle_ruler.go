package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRuler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="ipOTriangleRuler0" d="M13 35h12L13 23v12Z" clip-rule="evenodd"/><path id="ipOTriangleRuler1" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 35h12L13 23v12Z" clip-rule="evenodd"/></defs><g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 44L4 4v40h40Z"/><use href="#ipOTriangleRuler0" clip-rule="evenodd"/><use href="#ipOTriangleRuler0" clip-rule="evenodd"/><use href="#ipOTriangleRuler0" clip-rule="evenodd"/><use href="#ipOTriangleRuler0" clip-rule="evenodd"/><use href="#ipOTriangleRuler1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"/><use href="#ipOTriangleRuler1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"/><use href="#ipOTriangleRuler1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"/><use href="#ipOTriangleRuler1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 44v-3m-6 3v-3m-6 3v-3m-6 3v-3m-8-5h3m-3-6h3m-3-6h3m-3-6h3"/></g>`),
		g.Group(children),
	)
}