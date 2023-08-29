package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1024q0 69 15 136t44 130t70 118t95 100l-91 91q-62-55-110-120t-82-139t-51-153t-18-163q0-106 27-204t78-183t120-156t155-120t184-77t204-28h646L1251 93l90-90l317 317l-317 317l-90-90l163-163H768q-88 0-170 23t-153 64t-129 100t-100 130t-65 153t-23 170zm1659-574q62 54 110 119t82 139t51 154t18 162q0 106-27 204t-78 183t-120 156t-155 120t-184 77t-204 28H634l163 163l-90 90l-317-317l317-317l90 90l-163 163h646q88 0 170-23t153-64t129-100t100-130t65-153t23-170q0-70-15-137t-43-130t-70-117t-95-101l90-89z"/>`),
		g.Group(children),
	)
}