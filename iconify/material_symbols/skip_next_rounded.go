package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipNextRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 18q-.425 0-.713-.288T16.5 17V7q0-.425.288-.713T17.5 6q.425 0 .713.288T18.5 7v10q0 .425-.288.713T17.5 18ZM7.05 16.975q-.5.35-1.025.05t-.525-.9v-8.25q0-.6.525-.888t1.025.038l6.2 4.15q.45.3.45.825t-.45.825l-6.2 4.15Z"/>`),
		g.Group(children),
	)
}