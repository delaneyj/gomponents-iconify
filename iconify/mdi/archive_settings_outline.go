package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveSettingsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2v6h18V2H3m16 4H5V4h14v2m-1 3h2v11H4V9h2v9h12V9m-3 1.5V12H9v-1.5c0-.28.22-.5.5-.5h5c.28 0 .5.22.5.5M7 22h2v2H7v-2m4 0h2v2h-2v-2m4 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}