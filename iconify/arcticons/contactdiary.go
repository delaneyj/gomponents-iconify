package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contactdiary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="33" height="36.47" x="7.5" y="7.03" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.97"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.41 7.03V4.5m19.18 2.53V4.5M7.5 16.95h33m-25.48 8.93h18.83m-18.83 7.57h12.99"/>`),
		g.Group(children),
	)
}