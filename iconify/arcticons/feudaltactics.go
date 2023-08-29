package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feudaltactics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 21.36l9.15 5.28M14.842 5.5l-9.15 5.287V21.36l9.15 5.278v10.574L24 42.5l9.15-5.287V26.64l9.159-5.278V10.787L33.15 5.5L24 10.787L14.842 5.5Zm0 21.14L24 21.36m0 .001V10.787"/>`),
		g.Group(children),
	)
}