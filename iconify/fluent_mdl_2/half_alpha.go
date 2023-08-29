package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HalfAlpha(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1105 128l640 1920h-135l-171-512H615l-170 512H310L950 128h155zM658 1408h739L1027 300L658 1408z"/>`),
		g.Group(children),
	)
}