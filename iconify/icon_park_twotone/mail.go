package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMail0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 39h40V9H4v30Z"/><path stroke-linecap="round" d="m4 9l20 15L44 9"/><path stroke-linecap="round" d="M24 9H4v15m40 0V9H24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMail0)"/>`),
		g.Group(children),
	)
}