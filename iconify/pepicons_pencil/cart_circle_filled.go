package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CartCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCartCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M6.712 5.5H5.5a.5.5 0 0 1 0-1h1.603a.5.5 0 0 1 .485.379l1.897 7.6a.5.5 0 0 1-.97.242L6.712 5.5Z"/><path fill-rule="evenodd" d="M18.495 10.5h-7.99c-.15 0-.3.017-.447.05a2.02 2.02 0 0 0-1.508 2.418l.783 3.461A2.008 2.008 0 0 0 11.29 18h6.423a2.01 2.01 0 0 0 1.955-1.57l.783-3.462a2.012 2.012 0 0 0-1.955-2.468Zm-8.212 1.025a.992.992 0 0 1 .223-.025h7.989a1.013 1.013 0 0 1 .98 1.247l-.784 3.462a1.009 1.009 0 0 1-.98.791H11.29c-.468 0-.875-.328-.98-.791l-.783-3.462a1.02 1.02 0 0 1 .757-1.222Z" clip-rule="evenodd"/><path d="M20 19.75a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0Zm-7 0a1.75 1.75 0 1 1-3.5 0a1.75 1.75 0 0 1 3.5 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCartCircleFilled0)"/></g>`),
		g.Group(children),
	)
}