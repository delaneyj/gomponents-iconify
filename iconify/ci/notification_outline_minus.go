package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationOutlineMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22a2 2 0 0 1-2-2h4a2 2 0 0 1-2 2Zm8-3H4v-2l2-1v-5.5c0-3.462 1.421-5.707 4-6.32V2h4v2.18h.041l.042.01C16.611 4.843 18 7.08 18 10.5V16l2 1v2ZM12 5.75A3.6 3.6 0 0 0 8.875 7.2A5.692 5.692 0 0 0 8 10.5V17h8v-6.5a5.693 5.693 0 0 0-.875-3.3A3.6 3.6 0 0 0 12 5.75ZM15 13H9v-2h6v2Z"/>`),
		g.Group(children),
	)
}