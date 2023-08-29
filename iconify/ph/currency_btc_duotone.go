package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyBtcDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 160a40 40 0 0 1-40 40H80V48h60a36 36 0 0 1 0 72h12a40 40 0 0 1 40 40Z" opacity=".2"/><path d="M170.48 115.7A44 44 0 0 0 144 40.19V24a8 8 0 0 0-16 0v16h-16V24a8 8 0 0 0-16 0v16H64a8 8 0 0 0 0 16h8v136h-8a8 8 0 0 0 0 16h32v16a8 8 0 0 0 16 0v-16h16v16a8 8 0 0 0 16 0v-16h8a48 48 0 0 0 18.48-92.3ZM88 56h52a28 28 0 0 1 0 56H88Zm64 136H88v-64h64a32 32 0 0 1 0 64Z"/></g>`),
		g.Group(children),
	)
}