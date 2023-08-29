package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Felleskatalogen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.884" y="5.87" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.2 12.188h30.684V40.87c0 1.104-.897 2-2 2H12.2V12.188h0Zm5.476 3.927V39.34"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.408 39.34L22.62 27.727l14.788-11.533"/>`),
		g.Group(children),
	)
}