package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopHouseCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M4 14c-.92 0-1.352-1.137-.664-1.747l9-8a1 1 0 0 1 1.328 0l9 8c.688.61.255 1.747-.664 1.747h-1v7a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-7H4Zm6 6v-5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v5h3v-7a1 1 0 0 1 .512-.873L13 6.337l-6.512 5.79A1 1 0 0 1 7 13v7h3Zm2 0v-4h2v4h-2Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopHouseCircleFilled0)"/></g>`),
		g.Group(children),
	)
}