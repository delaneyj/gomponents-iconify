package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strava(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m20.516 23.927l-2.786-5.49h-4.083L20.517 32l6.865-13.563h-4.083zm-6.563-12.953l3.781 7.464h5.563L13.953 0L4.62 18.438h5.557z"/>`),
		g.Group(children),
	)
}