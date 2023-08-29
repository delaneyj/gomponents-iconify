package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoviePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 15h15l-.083-14H1v14zM11 2h1v1h-1V2zm3 2v8H3V4h11zM8 2h1v1H8V2zM5 2h1v1H5V2zM2 2h1v1H2V2zm.979 12h-1v-1h1v1zM6 14H5v-1h1v1zm3 0H8v-1h1v1zm3 0h-1v-1h1v1zm3 0h-1v-1h1v1zm0-11h-1V2h1v1z"/><path d="M7.003 9.865L7 6l3.949 2.015l-3.946 1.85z"/></g>`),
		g.Group(children),
	)
}