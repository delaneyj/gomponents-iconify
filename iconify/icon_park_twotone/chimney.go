package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chimney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChimney0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 44h40M13 14h22"/><path fill="#555" stroke-linejoin="round" d="M16 14h16l4 30H12l4-30Z"/><path stroke-linecap="round" d="M15 24h18M13 34h22"/><path stroke-linejoin="round" d="m32 14l4 30M16 14l-4 30"/><path stroke-linecap="round" stroke-linejoin="round" d="m24 8l.828-.828A4 4 0 0 1 27.657 6h.686a4 4 0 0 0 2.829-1.172L32 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChimney0)"/>`),
		g.Group(children),
	)
}