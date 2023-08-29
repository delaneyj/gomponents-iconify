package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JumpOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.414 21L3 1.586L1.586 3L3 4.414V21h16.586L21 22.414L22.414 21Zm-4.828-2H5V6.414L10.586 12l-.502.502l1.417 1.417l.502-.502L17.586 19Zm3.42-1.581L21 14.499v-1.01h-2v1.924l2.006 2.006ZM21 2.999h-8v2h4.586l-4.5 4.5l1.414 1.415l4.5-4.5V11h2V3Zm-10.49.003H9.5l-2.92-.006l2.007 2.006h1.923v-2Z"/>`),
		g.Group(children),
	)
}