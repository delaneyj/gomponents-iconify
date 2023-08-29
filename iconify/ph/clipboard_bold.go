package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 28h-34.53a51.88 51.88 0 0 0-74.94 0H56a20 20 0 0 0-20 20v168a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20V48a20 20 0 0 0-20-20Zm-44.29 32h-55.42a28 28 0 0 1 55.42 0ZM196 212H60V52h17.41A52.13 52.13 0 0 0 76 64v8a12 12 0 0 0 12 12h80a12 12 0 0 0 12-12v-8a52.13 52.13 0 0 0-1.41-12H196Z"/>`),
		g.Group(children),
	)
}