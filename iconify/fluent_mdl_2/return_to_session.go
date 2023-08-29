package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReturnToSession(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M885 512q-155 0-294 58T342 737Q233 846 175 985t-58 295q0 106 27 204t78 183t120 156t155 120t184 77t204 28h896v-128H885q-88 0-170-23t-153-64t-129-100t-100-130t-65-153t-23-170q0-88 23-170t64-153t100-129t130-100t153-65t170-23h821l-426 427l74 74l566-565l-566-565l-74 74l426 427H885z"/>`),
		g.Group(children),
	)
}