package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRuler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="ipSTriangleRuler0" fill="#fff" fill-rule="evenodd" d="M13 35h12L13 23v12Z" clip-rule="evenodd"/><path id="ipSTriangleRuler1" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 35h12L13 23v12Z" clip-rule="evenodd"/></defs><mask id="ipSTriangleRuler2"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 44L4 4v40h40Z"/><use href="#ipSTriangleRuler0" fill-rule="evenodd" clip-rule="evenodd"/><use href="#ipSTriangleRuler0" fill-rule="evenodd" clip-rule="evenodd"/><use href="#ipSTriangleRuler0" fill-rule="evenodd" clip-rule="evenodd"/><use href="#ipSTriangleRuler0" fill-rule="evenodd" clip-rule="evenodd"/><use href="#ipSTriangleRuler1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"/><use href="#ipSTriangleRuler1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"/><use href="#ipSTriangleRuler1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"/><use href="#ipSTriangleRuler1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 44v-3m-6 3v-3m-6 3v-3m-6 3v-3m-8-5h3m-3-6h3m-3-6h3m-3-6h3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTriangleRuler2)"/>`),
		g.Group(children),
	)
}