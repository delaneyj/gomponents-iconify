package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodepenFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#akarIconsCodepenFill0)"><path fill="currentColor" fill-rule="evenodd" d="M11.372.19c.38-.253.875-.253 1.256 0L23.492 7.4c.317.21.508.565.508.946v7.308c0 .38-.19.736-.508.947l-10.864 7.21c-.38.252-.875.252-1.256 0L.508 16.6A1.136 1.136 0 0 1 0 15.654V8.346c0-.38.19-.736.508-.947L11.372.19Zm-9.1 10.273v3.058l2.288-1.54l-2.288-1.518Zm4.337 2.878L3.18 15.648l7.684 5.1v-4.583L6.61 13.341Zm6.527 2.824v4.582l7.684-5.1l-3.43-2.306l-4.254 2.824Zm6.303-4.183l2.29 1.54v-3.06l-2.29 1.52Zm1.371-3.636l-3.41 2.263l-4.264-2.868V3.253l7.674 5.093Zm-9.946-5.093V7.74l-4.263 2.868L3.19 8.346l7.674-5.093ZM12 9.715l-3.35 2.254L12 14.192l3.35-2.223L12 9.715Z" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsCodepenFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}