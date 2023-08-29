package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationDeactivated(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22a2 2 0 0 1-2-2h4a2 2 0 0 1-2 2Zm7.585-.1l-2.9-2.9H4v-2l2-1v-5.5a10.57 10.57 0 0 1 .182-2L2.615 4.929L4.03 3.515L21 20.487L19.585 21.9ZM18 14.659L8.28 4.938A4.946 4.946 0 0 1 10 4.18V2h4v2.18c2.58.613 4 2.858 4 6.32v4.16v-.001Z"/>`),
		g.Group(children),
	)
}