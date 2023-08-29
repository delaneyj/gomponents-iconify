package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adblockbrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41 5.5H7.49a2 2 0 0 0-2 2V41c0 1.4.7 1.5.7 1.5l36.3-36.4s.3-.6-1.5-.6m-4.4 25.1v-1.8l-1.8-3.4l-.1-2.4l-2.3-1.4l-1.6.5l-2.4-.8l-.1-.3l2.5-2.5l.4.8l3.4.7l3.7-1.3a14 14 0 0 1-.8 12.9Zm-16.2-1.8l2.3-2.3h2l1.5 2.4l1.4.1l-.2 1.1l-2 1.9l-.1 1.5l-1.9 1.3l-.1 1l-.5-1v-5.6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35 13.7l-4.4-4.4H17.89l-8.6 8.6v12.8l4.3 4.3c14.8 13.6 35-6.6 21.4-21.3Z"/>`),
		g.Group(children),
	)
}