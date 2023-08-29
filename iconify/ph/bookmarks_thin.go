package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarksThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 28H96a12 12 0 0 0-12 12v20H64a12 12 0 0 0-12 12v152a4 4 0 0 0 6.33 3.25L112 188.92l53.69 38.33a3.94 3.94 0 0 0 2.31.75a4.08 4.08 0 0 0 1.83-.44A4 4 0 0 0 172 224v-47.08l25.68 18.33a3.94 3.94 0 0 0 2.32.75a4.08 4.08 0 0 0 1.83-.44A4 4 0 0 0 204 192V40a12 12 0 0 0-12-12Zm-28 188.23l-49.68-35.48a4 4 0 0 0-4.65 0L60 216.23V72a4 4 0 0 1 4-4h96a4 4 0 0 1 4 4Zm32-32l-24-17.14V72a12 12 0 0 0-12-12H92V40a4 4 0 0 1 4-4h96a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}