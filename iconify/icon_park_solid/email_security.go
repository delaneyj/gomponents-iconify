package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEmailSecurity0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24V9H4v30h20"/><path fill="#fff" d="M30 30c0-1 5-3 5-3s5 2 5 3c0 8-5 10-5 10s-5-2-5-10Z"/><path d="m4 9l20 15L44 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEmailSecurity0)"/>`),
		g.Group(children),
	)
}