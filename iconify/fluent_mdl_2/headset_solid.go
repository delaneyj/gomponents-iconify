package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 0q106 0 204 27t183 78t156 120t120 155t77 184t28 204v544q0 40-28 68t-68 28h-288V768h256q0-88-23-170t-64-153t-100-129t-130-100t-153-65t-170-23q-88 0-170 23t-153 64t-129 100t-100 130t-65 153t-23 170h256v640H512v128q0 53 20 99t55 82t81 55t100 20v-48q0-33 23-56t57-24h352q33 0 56 23t24 57v224q0 33-23 56t-57 24H848q-33 0-56-23t-24-57v-48q-80 0-150-30t-122-82t-82-122t-30-150v-128h-32q-40 0-68-28t-28-68V768q0-106 27-204t78-183t120-156t155-120t184-77t204-28z"/>`),
		g.Group(children),
	)
}