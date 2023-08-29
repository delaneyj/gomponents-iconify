package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicYou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="18.538" cy="35.378" r="8.122" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.66 35.378V12.834M28.69 4.5h6.864a2.03 2.03 0 0 1 2.03 2.03v4.273a2.03 2.03 0 0 1-2.03 2.03H26.66h0V6.53a2.03 2.03 0 0 1 2.03-2.03Z"/>`),
		g.Group(children),
	)
}