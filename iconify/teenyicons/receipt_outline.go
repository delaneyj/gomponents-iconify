package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5.5V0a.5.5 0 0 0-.5.5h.5Zm12 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 14l-.224.447A.5.5 0 0 0 14 14.5h-.5Zm-2-1l.224-.447a.5.5 0 0 0-.448 0l.224.447Zm-2 1l-.224.447a.5.5 0 0 0 .448 0L9.5 14.5Zm-2-1l.224-.447a.5.5 0 0 0-.448 0l.224.447Zm-2 1l-.224.447a.5.5 0 0 0 .448 0L5.5 14.5Zm-4 0H1a.5.5 0 0 0 .724.447L1.5 14.5Zm2-1l.224-.447a.5.5 0 0 0-.448 0l.224.447ZM1.5 1h12V0h-12v1ZM13 .5v14h1V.5h-1Zm.724 13.553l-2-1l-.448.894l2 1l.448-.894Zm-2.448-1l-2 1l.448.894l2-1l-.448-.894Zm-1.552 1l-2-1l-.448.894l2 1l.448-.894Zm-2.448-1l-2 1l.448.894l2-1l-.448-.894ZM2 14.5V.5H1v14h1Zm3.724-.447l-2-1l-.448.894l2 1l.448-.894Zm-2.448-1l-2 1l.448.894l2-1l-.448-.894ZM4 5h2V4H4v1Zm4 0h3V4H8v1ZM4 8h2V7H4v1Zm4 0h3V7H8v1Zm0 3h3v-1H8v1Z"/>`),
		g.Group(children),
	)
}