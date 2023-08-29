package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smallpdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 16.5h15v15h-15zm15 0V3.845M16.5 16.5V3.845m15 27.655h12.655M31.5 16.5h12.655M16.5 31.5v12.655m15-12.655v12.655M16.5 16.5H3.845m12.655 15H3.845"/>`),
		g.Group(children),
	)
}