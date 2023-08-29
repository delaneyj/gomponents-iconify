package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightWeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1457 1317q0 111-37 199t-104 149t-157 94t-199 33H512V128h438q89 0 167 24t136 73t91 121t34 166q0 72-19 136t-57 116t-92 90t-124 61v4q84 9 152 40t117 82t75 120t27 156zM627 234v644h232q81 0 153-21t127-63t87-108t33-153q0-84-26-141t-74-93t-113-50t-142-15H627zm713 1092q0-104-40-171t-105-105t-150-53t-173-15H627v706h309q88 0 162-21t128-65t84-113t30-163z"/>`),
		g.Group(children),
	)
}