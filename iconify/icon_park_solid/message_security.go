package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMessageSecurity0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M25.5 37H21l-10 5v-5H4V7h40v11"/><path fill="#fff" d="M29 25.2c0-1.067 7-3.2 7-3.2s7 2.133 7 3.2c0 8.533-7 12.8-7 12.8s-7-4.267-7-12.8Z"/><path d="M12 15h6m-6 6h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMessageSecurity0)"/>`),
		g.Group(children),
	)
}