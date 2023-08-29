package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationUnreadLinesBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12Z" opacity=".5"/><path d="M7 16.75a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5H7Zm0-3.5a.75.75 0 0 0 0 1.5h9a.75.75 0 0 0 0-1.5H7ZM22 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g>`),
		g.Group(children),
	)
}