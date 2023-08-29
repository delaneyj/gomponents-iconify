package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yasm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m460.046 461.888l-39.82-63.309a8.423 8.423 0 0 0-2.24-2.374l-11.04-367.228C399.556-95.911-95.62 219.957 16.523 275.084l330.28 164.743l41.943 66.906c10.934 20.33 85.184-23.255 71.3-44.845z"/>`),
		g.Group(children),
	)
}