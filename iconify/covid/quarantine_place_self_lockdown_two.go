package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuarantinePlaceSelfLockdownTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 11.25a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5Zm1.5 12l.75-4.5h1.5V16.5a3.75 3.75 0 0 0-7.5 0v2.25h1.5l.75 4.5h3Z"/><path d="M22.873 8.949L20.5 7.161V1.25H16v2.518L12 .75L1.127 8.949c-.224.208-.36.495-.377.8v12.522a.98.98 0 0 0 .978.979h20.544a.98.98 0 0 0 .978-.979V9.75a1.186 1.186 0 0 0-.377-.801v0Z"/></g>`),
		g.Group(children),
	)
}