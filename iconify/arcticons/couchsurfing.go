package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Couchsurfing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.93 28.92C34.47 28.92 22 28 22 35.51a4.06 4.06 0 0 0 3.78 3.92c1.37 0 4.15-.29 4.15-2.81C30 30 15.63 25.8 15.63 15.44c0-4.92 5.41-6.87 8.68-6.87c3.08 0 8.49 1.95 8.49 8.33c0 6.13-7.33 12.6-16.49 12.6a74.14 74.14 0 0 1-13.39-1.27"/>`),
		g.Group(children),
	)
}