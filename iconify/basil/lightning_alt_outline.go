package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.614 2.31a.75.75 0 0 1 .456.69v6.998H18a.75.75 0 0 1 .653 1.12l-.492.87a35.748 35.748 0 0 1-7.05 8.842l-.796.725A.75.75 0 0 1 9.06 21v-6.939H5a.75.75 0 0 1-.653-1.119a35.801 35.801 0 0 1 6.675-8.773l1.778-1.71a.75.75 0 0 1 .814-.149Zm-7.33 10.251H9.81a.75.75 0 0 1 .75.75v5.983a34.247 34.247 0 0 0 6.153-7.796H13.32a.75.75 0 0 1-.75-.75V4.762l-.508.488a34.3 34.3 0 0 0-5.777 7.311Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}