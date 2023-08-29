package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uikit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.697 3.292l-4.109 2.489l4.738 2.696v7.077l-6.365 3.538l-6.258-3.538v-5.485L1.596 7.956V18l10.219 6l10.589-6V6.002l-4.707-2.71zm-1.904-.989L11.813 0L7.665 2.568l4.032 2.218l4.096-2.483z"/>`),
		g.Group(children),
	)
}