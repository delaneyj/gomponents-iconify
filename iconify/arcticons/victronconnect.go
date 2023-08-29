package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Victronconnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.8 28.414A11.65 11.65 0 1 1 13.107 17.17m7.155 1.211a11.65 11.65 0 1 1 15.553 12.157"/><path fill="none" stroke="currentColor" d="m15.153 11.836l19.614 25.739M14.765 27.319l7.24-5.809"/><circle cx="12.871" cy="28.686" r="2.288" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}