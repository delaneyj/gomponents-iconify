package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiReader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="16.137" height="13.75" x="5.5" y="17.125" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".992"/><rect width="16.137" height="13.75" x="26.363" y="17.125" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".992"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.637 23.045h4.726m4.917-1.075v4.535M10.107 21.97v4.535"/>`),
		g.Group(children),
	)
}