package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpDownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.213 10.787c-11.715-11.715-30.711-11.715-42.426 0c-11.716 11.715-11.716 30.71 0 42.426c11.715 11.717 30.711 11.717 42.426 0c11.716-11.716 11.716-30.712 0-42.426zM48 36.487L32 51.999L16 36.487h11.372v-8.974H16L32 12l16 15.514H36.628v8.974H48z"/>`),
		g.Group(children),
	)
}