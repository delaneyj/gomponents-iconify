package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMusic0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M30 34.5a3.5 3.5 0 0 1 3.5-3.5H41v3.4a3.6 3.6 0 0 1-3.6 3.6h-3.9a3.5 3.5 0 0 1-3.5-3.5Zm-24 4A3.5 3.5 0 0 1 9.5 35H16v3.4a3.6 3.6 0 0 1-3.6 3.6H9.5A3.5 3.5 0 0 1 6 38.5Z"/><path stroke-linecap="round" d="M16 18.044v0l25-5.919M16 38V10l25-6v29.692"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMusic0)"/>`),
		g.Group(children),
	)
}