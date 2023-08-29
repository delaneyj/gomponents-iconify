package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.6 4h6.178a.6.6 0 0 1 .39.144l3.164 2.712a.6.6 0 0 0 .39.144H21.4a.6.6 0 0 1 .6.6v2.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6ZM22 10v4"/><path d="M2 10v9.4a.6.6 0 0 0 .6.6H13m6 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-dasharray=".3 2" d="M19 22a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`),
		g.Group(children),
	)
}