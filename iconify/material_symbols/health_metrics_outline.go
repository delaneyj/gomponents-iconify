package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthMetricsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22v-5H2V7h5V2h10v5h5v10h-5v5H7ZM4 11h5q.25 0 .475.125t.35.325l.875 1.3l1.35-4.05q.1-.3.362-.5T13 8q.25 0 .475.125t.35.325l1.7 2.55H20V9h-5V4H9v5H4v2Zm5 9h6v-5h5v-2h-5q-.25 0-.475-.125t-.375-.325l-.85-1.3l-1.35 4.05q-.1.3-.375.5t-.6.2q-.25 0-.475-.125t-.35-.325L8.45 13H4v2h5v5Zm3-8Z"/>`),
		g.Group(children),
	)
}