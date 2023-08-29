package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill-rule="evenodd" d="M8.5 19.5a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm6 0a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20 7.5h-8a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2Zm-9 2a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1v-7Z" clip-rule="evenodd"/><path d="M13.25 11.5a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Zm0 2a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Zm0 2a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Z"/><path fill-rule="evenodd" d="M9 10.5H5.71a2 2 0 0 0-1.91 1.41L3 14.5v2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2Zm-5 6v-1.85l.755-2.445a1 1 0 0 1 .955-.705H9a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.5 12.6H5.79a.5.5 0 0 0-.484.378l-.29 1.15L5 14.25v.85a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5Zm-1.5 2v-.288l.18-.712H7v1H6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}