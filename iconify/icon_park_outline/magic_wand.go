package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagicWand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7.58 35.42A19.908 19.908 0 0 1 4 24C4 12.954 12.954 4 24 4s20 8.954 20 20s-8.954 20-20 20a19.908 19.908 0 0 1-11.42-3.58m-5-5a20.114 20.114 0 0 0 5 5m-5-5L16 27m-3.42 13.42L21 32m-5-5l4-4l5 5l-4 4m-5-5l5 5m-4-18h4m-2-2v4m9 1h6m-3-3v6m1 9h4m-2-2v4"/>`),
		g.Group(children),
	)
}