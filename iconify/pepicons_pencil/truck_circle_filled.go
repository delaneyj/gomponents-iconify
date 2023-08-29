package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTruckCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M8.5 19.5a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm6 0a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20 7.5h-8a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2Zm-9 2a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1v-7Z" clip-rule="evenodd"/><path d="M13.25 11.5a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Zm0 2a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Zm0 2a.5.5 0 0 1 0-1h5.5a.5.5 0 0 1 0 1h-5.5Z"/><path fill-rule="evenodd" d="M9 10.5H5.71a2 2 0 0 0-1.91 1.41L3 14.5v2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2Zm-5 6v-1.85l.755-2.445a1 1 0 0 1 .955-.705H9a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.5 12.6H5.79a.5.5 0 0 0-.484.378l-.29 1.15L5 14.25v.85a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5Zm-1.5 2v-.288l.18-.712H7v1H6Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTruckCircleFilled0)"/></g>`),
		g.Group(children),
	)
}