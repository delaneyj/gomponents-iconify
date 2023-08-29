package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenSplitV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M2.5 12.5A2.5 2.5 0 0 0 0 15v70a2.5 2.5 0 0 0 2.5 2.5h95A2.5 2.5 0 0 0 100 85V15a2.5 2.5 0 0 0-2.5-2.5zm2.5 5h90v65H5Z" color="currentColor"/><path fill="currentColor" d="M50 20L35.004 34.395h9.869V43H9v4.5h82V43H55.125v-8.605h9.873L50 20zM9 52.5V57h35.873v8.604H35L50 80l15-14.396h-9.877V57H91v-4.5H9z" color="currentColor"/>`),
		g.Group(children),
	)
}