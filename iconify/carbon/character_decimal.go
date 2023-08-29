package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CharacterDecimal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 15h2v2h-2z"/><path fill="currentColor" d="M24 23h-4a2.002 2.002 0 0 1-2-2V11a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v10a2.003 2.003 0 0 1-2 2zm-4-12v10h4V11zm-9 4h2v2h-2z"/><path fill="currentColor" d="M14 23h-4a2.002 2.002 0 0 1-2-2V11a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v10a2.003 2.003 0 0 1-2 2zm-4-12v10h4V11zM4 21h2v2H4z"/>`),
		g.Group(children),
	)
}