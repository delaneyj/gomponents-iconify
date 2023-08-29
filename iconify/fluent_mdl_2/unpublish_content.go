package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnpublishContent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M768 347L365 749l-90-90l557-557l557 557l-90 90l-403-402v1317H768V347zm1280 1253q0 93-35 174t-96 143t-142 96t-175 35q-93 0-174-35t-143-96t-96-142t-35-175q0-93 35-174t96-143t142-96t175-35q93 0 174 35t143 96t96 142t35 175zm-272 267l-443-443q-26 39-39 84t-14 92q0 67 25 125t68 101t102 69t125 25q47 0 92-13t84-40zm144-267q0-66-25-124t-69-101t-102-69t-124-26q-47 0-92 13t-84 40l443 443q26-39 39-84t14-92zm-835 320q22 37 48 69t59 59H0v-384h128v256h957z"/>`),
		g.Group(children),
	)
}