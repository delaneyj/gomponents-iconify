package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataVisFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 3H5a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h22a2.003 2.003 0 0 0 2-2V5a2.002 2.002 0 0 0-2-2Zm0 6H17V5h10ZM15 27h-4v-4h4Zm0-6h-4v-4h4Zm-6 0H5v-4h4Zm2-6v-4h10v4Zm0-6V5h4v4Zm12 2h4v4h-4ZM9 5v10H5V5ZM5 23h4v4H5Zm12 4V17h10.001l.001 10Z"/>`),
		g.Group(children),
	)
}