package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vibrationstest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.17 38.17a21.5 21.5 0 0 0 0-28.37m-32.34.03a21.5 21.5 0 0 0 0 28.34M15.3 17.3a11 11 0 0 0 0 13.4m17.4 0a11 11 0 0 0 0-13.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.63 13.63a16.15 16.15 0 0 0 0 20.74m24.74 0a16.15 16.15 0 0 0 0-20.74"/>`),
		g.Group(children),
	)
}