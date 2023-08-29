package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBigDownLinesFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m9 8l-.117.007A1 1 0 0 0 8 9v1.999L5.414 11A2 2 0 0 0 4 14.414L10.586 21a2 2 0 0 0 2.828 0L20 14.414a2 2 0 0 0 .434-2.18l-.068-.145A2 2 0 0 0 18.586 11L16 10.999V9a1 1 0 0 0-1-1H9zm6-6a1 1 0 0 1 .117 1.993L15 4H9a1 1 0 0 1-.117-1.993L9 2h6zm0 3a1 1 0 0 1 .117 1.993L15 7H9a1 1 0 0 1-.117-1.993L9 5h6z"/></g>`),
		g.Group(children),
	)
}