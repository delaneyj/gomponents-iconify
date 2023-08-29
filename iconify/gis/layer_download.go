package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="m89.91 78.264l-9.178.007l8.211 9.67l-77.875.053l8.22-9.682l-9.188.008L.832 89.234A3.5 3.5 0 0 0 3.502 95l93-.064a3.5 3.5 0 0 0 2.666-5.766zM41.34 5v20.705H24.637L50 51.174l25.363-25.469H58.66V5H41.34zm30.912 32.324c-2.317 2.328-4.632 4.658-6.951 6.985h5l18.64 21.957l-77.873.052l18.686-22.01H34.7c-2.317-2.328-4.637-4.65-6.955-6.978a3.5 3.5 0 0 0-2.28 1.213L.833 67.559a3.5 3.5 0 0 0 2.67 5.765l93-.064a3.5 3.5 0 0 0 2.666-5.766L74.59 38.543a3.5 3.5 0 0 0-2.338-1.219z" color="currentColor"/>`),
		g.Group(children),
	)
}