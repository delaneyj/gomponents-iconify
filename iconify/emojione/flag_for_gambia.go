package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGambia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f42f4c" d="M32 2C19.3 2 8.5 9.9 4.1 21h55.8C55.5 9.9 44.7 2 32 2z"/><path fill="#699635" d="M32 62c12.7 0 23.5-7.9 27.9-19H4.1C8.5 54.1 19.3 62 32 62z"/><path fill="#2a5f9e" d="M62 32c0-3.1-.5-6.2-1.4-9H3.4C2.5 25.8 2 28.9 2 32s.5 6.2 1.4 9h57.2c.9-2.8 1.4-5.9 1.4-9"/><path fill="#fff" d="M60.3 22c-.1-.3-.3-.7-.4-1H4.1c-.1.3-.3.7-.4 1c-.1.3-.2.7-.3 1h57.2c-.1-.3-.2-.7-.3-1M3.7 42c.1.3.3.7.4 1h55.8c.1-.3.3-.7.4-1c.1-.3.2-.7.3-1H3.4c.1.3.2.7.3 1"/>`),
		g.Group(children),
	)
}