package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookBookmarkBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 20H72a36 36 0 0 0-36 36v168a12 12 0 0 0 12 12h144a12 12 0 0 0 0-24H60v-4a12 12 0 0 1 12-12h136a12 12 0 0 0 12-12V32a12 12 0 0 0-12-12Zm-88 24h40v60l-12.81-9.6a12 12 0 0 0-14.4 0L120 104Zm76 128H72a35.59 35.59 0 0 0-12 2.06V56a12 12 0 0 1 12-12h24v84a12 12 0 0 0 19.2 9.6L140 119l24.81 18.6A12 12 0 0 0 184 128V44h12Z"/>`),
		g.Group(children),
	)
}