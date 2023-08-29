package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><rect width="20" height="4" x="2" y="3" fill="currentColor" rx="2"/><path fill="currentColor" d="M19 9c.513 0 .936.463.993 1.06l.007.14v7.2c0 1.917-1.249 3.484-2.824 3.594L17 21H7c-1.598 0-2.904-1.499-2.995-3.388L4 17.4v-7.2C4 9.537 4.448 9 5 9h14zm-5 2h-4l-.117.007a1 1 0 0 0 0 1.986L10 13h4l.117-.007a1 1 0 0 0 0-1.986L14 11z"/></g>`),
		g.Group(children),
	)
}