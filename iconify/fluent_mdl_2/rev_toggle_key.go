package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RevToggleKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1152 512q155 0 294 58t249 167q109 109 167 248t58 295q0 106-27 204t-78 183t-120 156t-155 120t-184 77t-204 28H256v-128h896q88 0 170-23t153-64t129-100t100-130t65-153t23-170q0-88-23-170t-64-153t-100-129t-130-100t-153-65t-170-23H331l426 427l-74 74l-566-565L683 11l74 74l-426 427h821z"/>`),
		g.Group(children),
	)
}