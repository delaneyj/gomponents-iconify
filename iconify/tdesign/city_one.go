package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h10v8h10v12H2V2Zm10 10v8h4v-5h4v-3h-8Zm8 5h-2v3h2v-3Zm-10 3V4H4v16h6ZM8.004 6v2.004H6v-2h.004V6h2Zm0 5v2.004H6v-2h.004V11h2Zm0 5v2.004H6v-2h.004V16h2Z"/>`),
		g.Group(children),
	)
}