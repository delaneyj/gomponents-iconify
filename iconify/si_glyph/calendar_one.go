package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M13 .031V1h-2V.031H5V1H3V.031H0V16h16V.031h-3zm1.029 13.977H1.955v-9.07h12.074v9.07z"/><path d="M6.027 7.957h-.98v-.941h1.902v4.938h-.922V7.957zm4 0H9v-.941h1.953v4.938h-.926V7.957z"/></g>`),
		g.Group(children),
	)
}