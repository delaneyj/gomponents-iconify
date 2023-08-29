package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wirsindhandwerk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M50.772 479.812h83.36V367.847l-83.36 47.01Zm329.046 0h82.35v-64.956l-82.35-47.009Zm.006-448v219.756l-123.648-72.382l-121.672 72.382V31.812H50.772v360.794l205.404-122.287l205.993 122.287V31.812Z"/>`),
		g.Group(children),
	)
}