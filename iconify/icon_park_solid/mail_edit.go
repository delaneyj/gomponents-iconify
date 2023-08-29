package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMailEdit0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24V9H4v30h20"/><path fill="#fff" d="m35 39l8-7l-4-4l-8 7v4h4Z"/><path d="m4 9l20 15L44 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMailEdit0)"/>`),
		g.Group(children),
	)
}