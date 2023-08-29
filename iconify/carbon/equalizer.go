package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 28h8v2H2zm0-4h8v2H2zm10 4h8v2h-8zm0-4h8v2h-8zM2 20h8v2H2zm10 0h8v2h-8zM2 16h8v2H2zm10 0h8v2h-8zm0-4h8v2h-8zm0-4h8v2h-8zm0-4h8v2h-8zm10 24h8v2h-8zm0-4h8v2h-8z"/>`),
		g.Group(children),
	)
}