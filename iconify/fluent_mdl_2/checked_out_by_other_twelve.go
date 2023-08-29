package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckedOutByOtherTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M939 171q129 0 249 33t224 95t190 146t147 190t94 225t34 249q0 129-33 249t-95 224t-146 191t-190 147t-225 94t-249 34q-130 0-250-33t-224-95t-190-147t-147-190t-94-224t-34-250q0-129 33-249t95-224t147-190t190-147t224-94t250-34zm0 1740q111 0 213-28t192-81t162-126t125-162t81-191t29-214q0-110-28-212t-81-192t-126-162t-163-126t-191-81t-213-29q-111 0-213 28t-192 81t-162 126t-125 162t-81 191t-29 214q0 111 28 213t81 192t125 162t163 126t191 80t214 29zm256-1058h170v683H683v-171h391L537 828l121-120l537 537V853z"/>`),
		g.Group(children),
	)
}