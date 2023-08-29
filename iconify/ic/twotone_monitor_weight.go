package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneMonitorWeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 8.5h1v1h-1zm-3 0h1v1h-1zm1.5 0h1v1h-1z" opacity=".3"/><path fill="currentColor" d="M5 19h14V5H5v14zm7-13c1.66 0 3 1.34 3 3s-1.34 3-3 3s-3-1.34-3-3s1.34-3 3-3z" opacity=".3"/><path fill="currentColor" d="M12 12c1.66 0 3-1.34 3-3s-1.34-3-3-3s-3 1.34-3 3s1.34 3 3 3zm1-3.5h1v1h-1v-1zm-1.5 0h1v1h-1v-1zm-1.5 0h1v1h-1v-1z"/><path fill="currentColor" d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V5h14v14z"/>`),
		g.Group(children),
	)
}