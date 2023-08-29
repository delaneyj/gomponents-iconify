package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeilingFanLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3v2h3v5.27c-.62.36-1 1.02-1 1.73v1h4v-1c0-.71-.38-1.37-1-1.73V5h3V3H8m-2 9c-2.21 0-4 .67-4 1.5S3.79 15 6 15s4-.67 4-1.5S8.21 12 6 12m12 0c-2.21 0-4 .67-4 1.5s1.79 1.5 4 1.5s4-.67 4-1.5s-1.79-1.5-4-1.5m-8 2v1c0 .72.38 1.38 1 1.73c.62.36 1.38.36 2 0c.62-.35 1-1.02 1-1.73v-1h-4m3 5v3h-2v-3h2m2.88-2.54L18 18.59L16.59 20l-2.12-2.12l1.41-1.42m-6.34 1.42L7.41 20L6 18.59l2.12-2.12l1.42 1.41"/>`),
		g.Group(children),
	)
}