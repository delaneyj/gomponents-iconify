package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 116h-13.68l-28.54-49.95a12 12 0 0 0-10.42-6h-24.65l-13-32.46A12 12 0 0 0 138.58 20h-21.16a12 12 0 0 0-11.15 7.54L93.29 60H68.64a12 12 0 0 0-10.42 6l-28.54 50H16a4 4 0 0 0 0 8h12v84a12 12 0 0 0 12 12h24a12 12 0 0 0 12-12v-20h104v20a12 12 0 0 0 12 12h24a12 12 0 0 0 12-12v-84h12a4 4 0 0 0 0-8ZM113.7 30.51a4 4 0 0 1 3.72-2.51h21.16a4 4 0 0 1 3.72 2.51L154.09 60h-52.18ZM65.17 70a4 4 0 0 1 3.47-2h118.72a4 4 0 0 1 3.47 2l26.28 46H38.89ZM68 208a4 4 0 0 1-4 4H40a4 4 0 0 1-4-4v-20h32Zm148 4h-24a4 4 0 0 1-4-4v-20h32v20a4 4 0 0 1-4 4Zm4-32H36v-56h184ZM60 152a4 4 0 0 1 4-4h16a4 4 0 0 1 0 8H64a4 4 0 0 1-4-4Zm112 0a4 4 0 0 1 4-4h16a4 4 0 0 1 0 8h-16a4 4 0 0 1-4-4Z"/>`),
		g.Group(children),
	)
}