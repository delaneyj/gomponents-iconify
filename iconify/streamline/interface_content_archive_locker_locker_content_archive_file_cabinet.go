package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceContentArchiveLockerLockerContentArchiveFileCabinet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5.5h11a1 1 0 0 1 1 1v10a2 2 0 0 1-2 2h-9a2 2 0 0 1-2-2v-10a1 1 0 0 1 1-1ZM.5 7h13"/><circle cx="7" cy="3.5" r=".5"/><circle cx="7" cy="10" r=".5"/></g>`),
		g.Group(children),
	)
}