package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 5v1H5v22h22V6h-4V5h-2v1H11V5zM7 8h2v1h2V8h10v1h2V8h2v2H7zm0 4h18v14H7zm4 6v2h10v-2z"/>`),
		g.Group(children),
	)
}