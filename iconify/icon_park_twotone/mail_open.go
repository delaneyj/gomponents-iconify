package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMailOpen0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M44 18v21.818C44 41.023 43.105 42 42 42H6c-1.105 0-2-.977-2-2.182V18l20 13l20-13Z"/><path stroke-linecap="round" d="M4 17.784L24 4l20 13.784"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMailOpen0)"/>`),
		g.Group(children),
	)
}