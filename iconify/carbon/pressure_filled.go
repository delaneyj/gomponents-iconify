package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PressureFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 30H10v-5H6l10-9l10 9h-4zm-6-14L6 7h4V2h12v5h4z"/>`),
		g.Group(children),
	)
}