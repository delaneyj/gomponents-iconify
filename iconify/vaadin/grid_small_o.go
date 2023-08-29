package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridSmallO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 4h4V0H0v4zm1-3h2v2H1V1zm-1 9h4V6H0v4zm1-3h2v2H1V7zm-1 9h4v-4H0v4zm1-3h2v2H1v-2zm5-9h4V0H6v4zm1-3h2v2H7V1zm-1 9h4V6H6v4zm1-3h2v2H7V7zm-1 9h4v-4H6v4zm1-3h2v2H7v-2zm5-13v4h4V0h-4zm3 3h-2V1h2v2zm-3 7h4V6h-4v4zm1-3h2v2h-2V7zm-1 9h4v-4h-4v4zm1-3h2v2h-2v-2z"/>`),
		g.Group(children),
	)
}