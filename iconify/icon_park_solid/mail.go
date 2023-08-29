package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMail0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 39h40V9H4v30Z"/><path stroke="#000" stroke-linecap="round" d="m4 9l20 15L44 9"/><path stroke="#fff" stroke-linecap="round" d="M24 9H4v15m40 0V9H24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMail0)"/>`),
		g.Group(children),
	)
}