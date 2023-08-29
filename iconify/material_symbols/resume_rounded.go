package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResumeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17V7q0-.425.288-.713T7 6q.425 0 .713.288T8 7v10q0 .425-.288.713T7 18q-.425 0-.713-.288T6 17Zm5.525.1q-.5.3-1.012 0T10 16.225v-8.45q0-.575.513-.875t1.012 0l7.05 4.25q.5.3.5.85t-.5.85l-7.05 4.25Z"/>`),
		g.Group(children),
	)
}