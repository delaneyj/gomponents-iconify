package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceLandlinePhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 3h-2a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm-4 1H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11"/><path d="M12 8H6v3h6zm0 6v.01M9 14v.01M6 14v.01M12 17v.01M9 17v.01M6 17v.01"/></g>`),
		g.Group(children),
	)
}