package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuteChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1600 1152q93 0 174 35t143 96t96 142t35 175q0 93-35 174t-96 143t-142 96t-175 35q-93 0-174-35t-143-96t-96-142t-35-175q0-93 35-174t96-143t142-96t175-35zm-320 448q0 66 25 124t68 102t102 69t125 25q47 0 92-13t84-40l-443-443q-26 39-39 84t-14 92zm587 176q26-39 39-84t14-92q0-66-25-124t-69-101t-102-69t-124-26q-47 0-92 13t-84 40l443 443zm181-1648v1089q-28-33-59-62t-69-52V256H128v1152h256v267l267-267h392q-22 63-29 128H704l-448 448v-448H0V128h2048z"/>`),
		g.Group(children),
	)
}