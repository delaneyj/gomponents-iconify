package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardSixtyNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="10.87" r="2.77" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 16.11v20.43a3.36 3.36 0 0 1-3.36 3.37h0A3.35 3.35 0 0 1 37 37.6l-6.28-18.83A3.89 3.89 0 0 0 27 16.11h-3v23.8h-3a3.9 3.9 0 0 1-3.69-2.67l-6.26-18.83a3.36 3.36 0 0 0-3.19-2.3h0a3.35 3.35 0 0 0-3.36 3.36V39.9"/>`),
		g.Group(children),
	)
}