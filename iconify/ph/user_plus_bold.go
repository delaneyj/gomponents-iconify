package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserPlusBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M256 136a12 12 0 0 1-12 12h-8v8a12 12 0 0 1-24 0v-8h-8a12 12 0 0 1 0-24h8v-8a12 12 0 0 1 24 0v8h8a12 12 0 0 1 12 12Zm-54.81 56.28a12 12 0 1 1-18.38 15.44C169.12 191.42 145 172 108 172c-28.89 0-55.46 12.68-74.81 35.72a12 12 0 0 1-18.38-15.44a124.08 124.08 0 0 1 48.69-35.75a72 72 0 1 1 89 0a124 124 0 0 1 48.69 35.75ZM108 148a48 48 0 1 0-48-48a48.05 48.05 0 0 0 48 48Z"/>`),
		g.Group(children),
	)
}