package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50 0a3.5 3.5 0 0 0-3.5 3.5V93h-9.57a3.5 3.5 0 0 0-3.5 3.5a3.5 3.5 0 0 0 3.5 3.5h25a3.5 3.5 0 0 0 3.5-3.5a3.5 3.5 0 0 0-3.5-3.5H53.5V47h43a3.5 3.5 0 0 0 3.5-3.5v-40A3.5 3.5 0 0 0 96.5 0h-46a3.5 3.5 0 0 0-.254.01A3.5 3.5 0 0 0 50 0z" color="currentColor"/>`),
		g.Group(children),
	)
}