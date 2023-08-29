package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMailReview0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M44 24v16.818C44 42.023 43.105 43 42 43H6c-1.105 0-2-.977-2-2.182V24l20 13l20-13Z"/><path stroke="#fff" stroke-linecap="round" d="m4 23.784l10-6.892m30 6.892l-10-6.892"/><path fill="#fff" stroke="#fff" d="M34 5H14v24.415a2 2 0 0 0 .91 1.677l8 5.2a2 2 0 0 0 2.18 0l8-5.2a2 2 0 0 0 .91-1.677V5Z"/><path stroke="#000" stroke-linecap="round" d="M20 13h4m-4 6h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMailReview0)"/>`),
		g.Group(children),
	)
}