package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkullLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 106a26 26 0 1 0 26 26a26 26 0 0 0-26-26Zm0 40a14 14 0 1 1 14-14a14 14 0 0 1-14 14Zm72-40a26 26 0 1 0 26 26a26 26 0 0 0-26-26Zm0 40a14 14 0 1 1 14-14a14 14 0 0 1-14 14ZM128 18C71.76 18 26 62 26 116c0 33.77 18.3 65.31 48 83.15V216a14 14 0 0 0 14 14h80a14 14 0 0 0 14-14v-16.85c29.7-17.84 48-49.38 48-83.15c0-54-45.76-98-102-98Zm45.09 172.44a6 6 0 0 0-3.09 5.25V216a2 2 0 0 1-2 2h-18v-26a6 6 0 0 0-12 0v26h-20v-26a6 6 0 0 0-12 0v26H88a2 2 0 0 1-2-2v-20.31a6 6 0 0 0-3.09-5.25C55.21 175.09 38 146.56 38 116c0-47.42 40.37-86 90-86s90 38.58 90 86c0 30.56-17.21 59.09-44.91 74.44Z"/>`),
		g.Group(children),
	)
}