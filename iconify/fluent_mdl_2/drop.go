package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1606 1166q58 114 58 242q0 88-23 170t-64 153t-100 129t-130 100t-153 65t-170 23q-88 0-170-23t-153-64t-129-100t-100-130t-65-153t-23-170q0-127 56-241L1019 0l587 1166z"/>`),
		g.Group(children),
	)
}