package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.5 5.5h17v17h-17zm0 20h17v17h-17zm-20-20h17v17h-17zm0 20h17v17h-17z"/>`),
		g.Group(children),
	)
}