package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilHouseCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M4 13h1.389v7a.5.5 0 0 0 .5.5H19.11a.5.5 0 0 0 .5-.5v-7H21a.5.5 0 0 0 .33-.875l-8.5-7.5a.5.5 0 0 0-.66 0l-8.5 7.5A.5.5 0 0 0 4 13Zm1.889-1h-.567L12.5 5.667L19.678 12h-.567a.5.5 0 0 0-.5.5v7H6.39v-7a.5.5 0 0 0-.5-.5Z"/><path d="M13.708 14.5h-2.5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h2.5a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Zm-2.5 5v-4h2.5v4h-2.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilHouseCircleFilled0)"/></g>`),
		g.Group(children),
	)
}