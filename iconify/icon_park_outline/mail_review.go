package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M44 24v16.818C44 42.023 43.105 43 42 43H6c-1.105 0-2-.977-2-2.182V24l20 13l20-13Z"/><path stroke-linecap="round" d="m4 23.784l10-6.892m30 6.892l-10-6.892"/><path d="M34 5H14v24.415a2 2 0 0 0 .91 1.677l8 5.2a2 2 0 0 0 2.18 0l8-5.2a2 2 0 0 0 .91-1.677V5Z"/><path stroke-linecap="round" d="M20 13h4m-4 6h8"/></g>`),
		g.Group(children),
	)
}