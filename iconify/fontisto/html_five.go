package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.95 7.035l.24-2.625H3.929l.705 8.01h9.18l-.33 3.42l-2.955.795l-2.94-.795l-.195-2.1H4.769l.33 4.17l5.43 1.5h.06v-.015l5.385-1.485l.75-8.16h-9.66l-.225-2.715zM0 0h21.12L19.2 21.57L10.53 24l-8.61-2.43z"/>`),
		g.Group(children),
	)
}