package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PublicContactCardMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 896h640V768H256v128zm640 256H512v128h384v-128zm326-80q-46 26-82 62t-62 79t-40 93t-14 102h128q0-53 20-99t55-82t81-55t100-20q53 0 99 20t82 55t55 81t20 100h128q0-52-14-101t-39-93t-62-80t-83-62q33-35 51-81t19-95q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100q0 49 18 95t52 81zm314-176q0 27-10 50t-27 40t-41 28t-50 10q-27 0-50-10t-40-27t-28-41t-10-50q0-27 10-50t27-40t41-28t50-10q27 0 50 10t40 27t28 41t10 50zM319 1411L2 1728l317 317l91-90l-163-163h658l-163 163l91 90l317-317l-317-317l-91 90l163 163H247l163-163l-91-90zM0 256v1216l128-128V384h1792v1280h-640v128h768V256H0z"/>`),
		g.Group(children),
	)
}