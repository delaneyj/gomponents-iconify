package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 96a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm92-40v144a20 20 0 0 1-20 20H40a20 20 0 0 1-20-20V56a20 20 0 0 1 20-20h176a20 20 0 0 1 20 20ZM44 60v79.72l33.86-33.86a20 20 0 0 1 28.28 0L147.31 147l17.18-17.17a20 20 0 0 1 28.28 0L212 149.09V60Zm0 136h118.34L92 125.66l-48 48Zm168 0v-13l-33.37-33.37L164.28 164l32 32Z"/>`),
		g.Group(children),
	)
}