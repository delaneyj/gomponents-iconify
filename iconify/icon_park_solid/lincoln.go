package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lincoln(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLincoln0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M30.19 4H17.81a2 2 0 0 0-1.99 1.801l-1.8 18a2.005 2.005 0 0 0 0 .398l1.8 18A2 2 0 0 0 17.81 44h12.38a2 2 0 0 0 1.99-1.801l1.8-18a2.005 2.005 0 0 0 0-.398l-1.8-18A2 2 0 0 0 30.19 4Z"/><path stroke="#000" stroke-linecap="round" d="M14 24h20M24 4v40"/><path stroke="#fff" stroke-linecap="round" d="M20 4h8m-8 40h8M15 14l-.98 9.801a2.005 2.005 0 0 0 0 .398L15 34m18-20l.98 9.801c.013.132.013.266 0 .398L33 34"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLincoln0)"/>`),
		g.Group(children),
	)
}