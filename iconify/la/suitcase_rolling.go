package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseRolling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 3c-1.654 0-3 1.346-3 3v3h2V6a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v3h2V6c0-1.654-1.346-3-3-3h-4zm-7 8v15h2v2h2v-2h10v2h2v-2h2V11H7zm2 2h14v11H9V13zm2 3v2h10v-2H11zm0 4v2h10v-2H11z"/>`),
		g.Group(children),
	)
}