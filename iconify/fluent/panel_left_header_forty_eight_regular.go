package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftHeaderFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 13.25A6.25 6.25 0 0 1 10.25 7h27.5A6.25 6.25 0 0 1 44 13.25v21.5A6.25 6.25 0 0 1 37.75 41h-27.5A6.25 6.25 0 0 1 4 34.75v-21.5Zm6.25-3.75a3.75 3.75 0 0 0-3.75 3.75v21.5a3.75 3.75 0 0 0 3.75 3.75h5.5v-29h-5.5Zm8 11v18h19.5a3.75 3.75 0 0 0 3.75-3.75V20.5H18.25Zm0-2.5H41.5v-4.75a3.75 3.75 0 0 0-3.75-3.75h-19.5V18Z"/>`),
		g.Group(children),
	)
}