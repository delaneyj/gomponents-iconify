package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 96a92 92 0 1 0-152 69.69V240a12 12 0 0 0 17.37 10.73L128 229.42l42.64 21.31A12 12 0 0 0 188 240v-74.31A91.86 91.86 0 0 0 220 96ZM60 96a68 68 0 1 1 68 68a68.07 68.07 0 0 1-68-68Zm104 124.59l-30.64-15.32a12 12 0 0 0-10.74 0L92 220.58v-39.92a92 92 0 0 0 72 0ZM128 148a52 52 0 1 0-52-52a52.06 52.06 0 0 0 52 52Zm0-80a28 28 0 1 1-28 28a28 28 0 0 1 28-28Z"/>`),
		g.Group(children),
	)
}