package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yandex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.083 14.8L2.985 24H-.001l4.5-9.834C2.385 13.092.974 11.148.974 7.552C.969 2.518 4.159.001 7.953.001h3.858v24H9.229v-9.2H7.083zM9.23 2.18H7.852c-2.08 0-4.097 1.378-4.097 5.372c0 3.858 1.847 5.1 4.097 5.1H9.23z"/>`),
		g.Group(children),
	)
}