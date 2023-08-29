package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionPlayCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd"><path d="M7.25 12v5a2 2 0 0 0 2 2h7.5a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2h-7.5a2 2 0 0 0-2 2Zm-1 0v5a3 3 0 0 0 3 3h7.5a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3h-7.5a3 3 0 0 0-3 3Z"/><path d="M14.5 14.5L12 16v-3l2.5 1.5Z"/><path d="M15 14.5a.5.5 0 0 1-.243.429l-2.5 1.5A.5.5 0 0 1 11.5 16v-3a.5.5 0 0 1 .757-.429l2.5 1.5A.5.5 0 0 1 15 14.5Zm-2.5-.617v1.234l1.028-.617l-1.028-.617Zm3.062-7.773a.5.5 0 0 1 .078.702l-2 2.5a.5.5 0 1 1-.78-.624l2-2.5a.5.5 0 0 1 .702-.078Z"/><path d="M9.938 6.11a.5.5 0 0 1 .702.078l2 2.5a.5.5 0 1 1-.78.624l-2-2.5a.5.5 0 0 1 .078-.702Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}