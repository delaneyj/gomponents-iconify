package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudHsm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 21.142V10h-5v2h3v2h-3v2h3v5.142c-1.72.447-3 2-3 3.858c0 2.206 1.794 4 4 4s4-1.794 4-4c0-1.859-1.28-3.41-3-3.858ZM27 27c-1.103 0-2-.897-2-2s.897-2 2-2s2 .897 2 2s-.897 2-2 2Z"/><circle cx="9" cy="24" r="1" fill="currentColor"/><circle cx="9" cy="8" r="1" fill="currentColor"/><path fill="currentColor" d="M26 6v2h2V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4v8H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h14v-2H6v-4h14v-2h-8v-8h8v-2H6V6h20Z"/>`),
		g.Group(children),
	)
}