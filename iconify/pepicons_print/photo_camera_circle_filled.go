package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCameraCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd"><path d="M9.172 8.232L8.762 9.5H7.5A2.5 2.5 0 0 0 5 12v6a2.5 2.5 0 0 0 2.5 2.5h11A2.5 2.5 0 0 0 21 18v-6a2.5 2.5 0 0 0-2.5-2.5h-1.263l-.409-1.268a2.5 2.5 0 0 0-2.38-1.732h-2.897a2.5 2.5 0 0 0-2.38 1.732ZM7.5 10.5h1.99l.633-1.96a1.5 1.5 0 0 1 1.428-1.04h2.898a1.5 1.5 0 0 1 1.427 1.04l.633 1.96H18.5A1.5 1.5 0 0 1 20 12v6a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 6 18v-6a1.5 1.5 0 0 1 1.5-1.5Z"/><path d="M10 14.5a3 3 0 1 0 6 0a3 3 0 0 0-6 0Zm5 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}