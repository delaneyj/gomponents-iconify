package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailChatBubbleTypingSquareMessagesMessageBubbleTypingSquareChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="4.25" cy="6.5" r=".5"/><circle cx="7.5" cy="6.5" r=".5"/><circle cx="10.75" cy="6.5" r=".5"/><path d="m4.5 12.5l-4 1l1-3v-9a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1Z"/></g>`),
		g.Group(children),
	)
}