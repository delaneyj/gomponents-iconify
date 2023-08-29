package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sneer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm5-3h4v2H6V9Zm8 0h4v2h-4V9Zm2.718 4.78l-.25.97c-.269 1.045-.793 1.895-1.613 2.467c-.806.563-1.792.783-2.855.783h-1v-2h1c.8 0 1.343-.167 1.71-.423c.353-.246.647-.646.822-1.326l.249-.969l1.937.499Z"/>`),
		g.Group(children),
	)
}