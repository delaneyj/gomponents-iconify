package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Color(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.5 14h-3c-.28 0-.5-.22-.5-.5v-3c0-.13.05-.26.15-.35l5-5c.2-.2.51-.2.71 0l3 3c.2.2.2.51 0 .71l-5 5a.51.51 0 0 1-.35.15ZM3 13h2.29l4.5-4.5L7.5 6.21L3 10.71V13Zm10.37-7.38L11.99 7l-3-3l1.38-1.38c.42-.42.96-.62 1.5-.62s1.08.2 1.5.62c.83.83.83 2.17 0 3Z"/><path fill="currentColor" d="M12.5 8a.47.47 0 0 1-.35-.15l-4-4c-.2-.2-.2-.51 0-.71c.2-.2.51-.2.71 0l4 4c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/>`),
		g.Group(children),
	)
}