package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3a1 1 0 0 0-.976.783l-2 9A1 1 0 0 0 2 13v7a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-7a.988.988 0 0 0-.024-.217l-2-9A1 1 0 0 0 19 3H5Zm14.753 9H15a3 3 0 1 1-6 0H4.247l1.555-7h12.396l1.555 7Z"/>`),
		g.Group(children),
	)
}