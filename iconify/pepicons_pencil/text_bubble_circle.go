package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBubbleCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g transform="translate(3 3)"><path fill-rule="evenodd" d="M1.5 12.102a2.5 2.5 0 0 0 2.5 2.5h.249c.307 1.238.741 1.898 1.455 1.898c.603 0 1.519-.62 2.94-1.898H16a2.5 2.5 0 0 0 2.5-2.5V6A2.5 2.5 0 0 0 16 3.5H4A2.5 2.5 0 0 0 1.5 6v6.102Zm1 0V6A1.5 1.5 0 0 1 4 4.5h12A1.5 1.5 0 0 1 17.5 6v6.102a1.5 1.5 0 0 1-1.5 1.5H8.257l-.143.13C6.834 14.898 5.96 15.5 5.704 15.5c-.092 0-.35-.463-.567-1.5l-.083-.398H4a1.5 1.5 0 0 1-1.5-1.5Z" clip-rule="evenodd"/><circle cx="7" cy="9" r="1"/><circle cx="10" cy="9" r="1"/><circle cx="13" cy="9" r="1"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}