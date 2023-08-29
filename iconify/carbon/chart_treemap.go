package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartTreemap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H4a2.002 2.002 0 0 0-2 2v24a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2Zm0 12h-5V4h5ZM16 4h5v10h-5Zm-2 0v16H4V4ZM4 22h10v6H4Zm12 6V16h12v12Z"/>`),
		g.Group(children),
	)
}