package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droidedit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.895 39.45h20.21"/><rect width="39" height="25.354" x="4.5" y="8.55" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 14.939h26M11 19.13h11.305m3.39 0H37m-26 4.192h5.746m3.203 0H37m0 4.192h-5.746m-3.203 0H11m17.333 6.389v5.547m-8.666-5.547v5.547"/>`),
		g.Group(children),
	)
}