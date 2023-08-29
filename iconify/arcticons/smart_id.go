package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartId(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.75 32.734V2.92A21.6 21.6 0 0 1 24 2.5c11.874 0 21.5 9.626 21.5 21.5c0 7.292-3.63 13.735-9.18 17.623"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.25 24.708V45.08h0a21.6 21.6 0 0 1-4.25.42C12.126 45.5 2.5 35.874 2.5 24c0-7.292 3.63-13.735 9.18-17.623"/><circle cx="28.25" cy="15.439" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}