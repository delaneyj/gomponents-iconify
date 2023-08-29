package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircledAnticlockwiseArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="26.68" fill="#fff" fill-rule="evenodd"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><circle cx="36" cy="36" r="26.68" stroke-width="4.74"/><path stroke-miterlimit="7" stroke-width="7" d="m20.66 30.58l5.421 5.421l5.424-5.424"/><path stroke-miterlimit="7" stroke-width="7" d="M25.13 31.65c0-6.003 4.869-10.87 10.87-10.87c5.999 0 10.87 4.865 10.87 10.87v8.693v-8.693v8.693c0 6.003-4.869 10.87-10.87 10.87a10.876 10.876 0 0 1-9.739-6.034" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}