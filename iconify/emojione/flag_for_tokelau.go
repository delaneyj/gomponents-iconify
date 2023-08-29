package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTokelau(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#2a5f9e" d="M53.4 53h-42s26-29.4 42.4-38.2c0 0-11.8 21.5 3 34.2C60 44.2 62 38.3 62 32C62 15.4 48.6 2 32 2S2 15.4 2 32c0 9.8 4.7 18.5 12 24h36c1.2-.9 2.4-1.9 3.4-3M32 62c4.7 0 9.1-1.1 13.1-3H18.9c4 1.9 8.4 3 13.1 3"/><path fill="#ffce31" d="M53.7 14.8C37.4 23.6 11.4 53 11.4 53h42c1.2-1.2 2.3-2.6 3.3-4c-14.8-12.7-3-34.2-3-34.2M14 56c1.5 1.1 3.2 2.2 4.9 3h26.2c1.7-.8 3.4-1.9 4.9-3H14"/><path fill="#fff" d="M18.5 10.3L19 12h1.8l-1.4 1l.5 1.7l-1.4-1l-1.4 1l.5-1.7l-1.4-1h1.7zm8.6 7.8l.5 1.5h1.6l-1.3.9l.5 1.5l-1.3-.9l-1.2.9l.5-1.5l-1.3-.9h1.6zM8.3 20.5l.5 1.7h1.8l-1.4 1l.5 1.7l-1.4-1l-1.4 1l.5-1.7l-1.4-1h1.8zm10.2 12.8l.6 2h2.1l-1.7 1.2l.6 2l-1.6-1.2l-1.7 1.2l.6-2l-1.7-1.2h2.1z"/>`),
		g.Group(children),
	)
}