package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailChatBubbleTypingOvalMessagesMessageBubbleTypingChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3.5" cy="7" r=".5"/><circle cx="6.75" cy="7" r=".5"/><circle cx="10" cy="7" r=".5"/><path d="M7 .5a6.5 6.5 0 0 0-5.41 10.1L.5 13.5l3.65-.66A6.5 6.5 0 1 0 7 .5Z"/></g>`),
		g.Group(children),
	)
}