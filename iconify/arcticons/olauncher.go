package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Olauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.203 39.203A21.433 21.433 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5a21.433 21.433 0 0 1 15.203 6.297"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.114 32.03C37.163 37.94 31.055 42 24 42c-9.941 0-18-8.059-18-18S14.059 6 24 6c7.055 0 13.162 4.06 16.114 9.97"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.872 32.327A14.483 14.483 0 0 1 24 38.5c-8.008 0-14.5-6.492-14.5-14.5S15.992 9.5 24 9.5c4.91 0 9.248 2.44 11.872 6.173"/>`),
		g.Group(children),
	)
}