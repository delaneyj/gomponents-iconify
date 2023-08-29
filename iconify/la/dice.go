package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 1L0 13l12 12l12-12L12 1zm0 2.8l9.2 9.2l-9.2 9.2L2.8 13L12 3.8zM12 8a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm-4 4a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm4 0a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm4 0a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm10.8 1l-2 2H30v12H18v-5.2l-2 2V29h16V13h-5.2zM12 16a1 1 0 0 0 0 2a1 1 0 0 0 0-2zm12 4a1 1 0 0 0 0 2a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}