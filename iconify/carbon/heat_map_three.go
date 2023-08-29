package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatMapThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 3H5a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h22a2.002 2.002 0 0 0 2-2V5a2.002 2.002 0 0 0-2-2Zm-8 6h-6V5h6Zm0 2v4h-6v-4Zm-8 0v4H5v-4Zm0 6v4H5v-4Zm2 0h6v4h-6Zm8-2v-4h6v4ZM5 23h6v4H5Zm16 4v-4h6v4Z"/>`),
		g.Group(children),
	)
}