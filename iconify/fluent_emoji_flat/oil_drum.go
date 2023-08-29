package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilDrum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M25 6H7v22h18V6Z"/><path fill="#212121" d="M23.02 2h-3.03c-.55 0-.99.44-.99.98V4h5V2.98c0-.54-.44-.98-.98-.98Z"/><path fill="#26EAFC" d="M20 6h1v22h-1V6Zm2 0h2v22h-2V6Z"/><path fill="#0074BA" d="M7 6h18c.55 0 1-.45 1-1s-.45-1-1-1H7c-.55 0-1 .45-1 1s.45 1 1 1Zm0 24h18c.55 0 1-.45 1-1s-.45-1-1-1H7c-.55 0-1 .45-1 1s.45 1 1 1Zm18-16H7c-.55 0-1-.45-1-1s.45-1 1-1h18c.55 0 1 .45 1 1s-.45 1-1 1ZM7 22h18c.55 0 1-.45 1-1s-.45-1-1-1H7c-.55 0-1 .45-1 1s.45 1 1 1Z"/></g>`),
		g.Group(children),
	)
}