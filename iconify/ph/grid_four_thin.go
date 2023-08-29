package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridFourThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 44H56a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12Zm4 12v68h-72V52h68a4 4 0 0 1 4 4ZM56 52h68v72H52V56a4 4 0 0 1 4-4Zm-4 148v-68h72v72H56a4 4 0 0 1-4-4Zm148 4h-68v-72h72v68a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}