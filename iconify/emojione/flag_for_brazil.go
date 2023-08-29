package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForBrazil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#699635" d="M61.5 26.8C59.1 12.7 46.8 2 32 2S4.9 12.7 2.5 26.8L32 12l29.5 14.8zm-59 10.4C4.9 51.3 17.2 62 32 62s27.1-10.7 29.5-24.8L32 52L2.5 37.2z"/><path fill="#ffe62e" d="M32 12L2.5 26.8C2.2 28.5 2 30.2 2 32s.2 3.5.5 5.2L32 52l29.5-14.8c.3-1.7.5-3.4.5-5.2s-.2-3.5-.5-5.2L32 12"/><g fill="#428bc1"><path d="M26 28.4c-3.2 0-6.2.7-8.9 1.9c-.1.6-.1 1.1-.1 1.7c0 8.3 6.7 15 15 15c5.6 0 10.5-3.1 13.1-7.6c-3.7-6.5-10.9-11-19.1-11"/><path d="M46.8 34.4c.1-.8.2-1.6.2-2.4c0-8.3-6.7-15-15-15c-5.9 0-11 3.4-13.5 8.4c2.4-.7 4.9-1.1 7.5-1.1c8.5 0 16 4 20.8 10.1"/></g><g fill="#fff"><path d="M26 24.3c-2.6 0-5.1.4-7.5 1.1c-.7 1.5-1.2 3.1-1.4 4.9c2.7-1.2 5.7-1.9 8.9-1.9c8.2 0 15.4 4.4 19.1 10.9c.9-1.5 1.4-3.2 1.7-4.9C42 28.3 34.5 24.3 26 24.3"/><circle cx="22" cy="32" r="1"/><circle cx="26" cy="38" r="1"/><circle cx="32" cy="38" r="1"/><circle cx="32" cy="42" r="1"/><circle cx="40" cy="38" r="1"/><circle cx="40" cy="42" r="1"/><circle cx="36" cy="40" r="1"/><circle cx="22" cy="36" r="1"/></g>`),
		g.Group(children),
	)
}