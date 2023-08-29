package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailUnpacking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M44 18v21.818C44 41.023 43.105 42 42 42H6c-1.105 0-2-.977-2-2.182V18l20 16l20-16Z"/><path stroke-linecap="round" d="M4 17.784L24 4l20 13.784"/><path d="M34 18H14v8l10 8l10-8v-8Z"/></g>`),
		g.Group(children),
	)
}