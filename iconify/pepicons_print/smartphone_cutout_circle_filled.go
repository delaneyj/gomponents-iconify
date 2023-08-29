package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneCutoutCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd"><path d="M7 5a1.5 1.5 0 0 1 1.5-1.5h9A1.5 1.5 0 0 1 19 5v16a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 7 21V5Zm1.5-.5A.5.5 0 0 0 8 5v16a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V5a.5.5 0 0 0-.5-.5h-9Z"/><path d="M10.5 6a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}