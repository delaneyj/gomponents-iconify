package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoodLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M23.955 30.007s-11.038-10.207.252-12.014c0 0 10.673 1.405-.252 12.014Z"/><path d="M14.763 4.784s25.583 2.029 27.943 29.61M4.92 14.261S5.911 39.5 34.162 42.94"/><path d="M4.814 33.535s.06-24.951 28.112-28.916M14.28 43.094s25.87-.627 28.684-28.714"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}