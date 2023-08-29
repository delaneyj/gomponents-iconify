package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessagePrivacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMessagePrivacy0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M25.5 37H21l-10 5v-5H4V7h40v11"/><path fill="#fff" d="M30 27h14v8H30z"/><path d="M40 27v-3a3 3 0 1 0-6 0v3M12 15h6m-6 6h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMessagePrivacy0)"/>`),
		g.Group(children),
	)
}