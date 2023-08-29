package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Giftbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 768v1280H128V768h292q-35-59-35-128q0-53 20-99t54-82t81-55t100-20q97 0 181 45t139 127q54-81 138-126t182-46q53 0 99 20t82 55t55 81t20 100q0 34-9 66t-27 62h292zM896 896H256v1024h640V896zm0-128q0-53-20-99t-55-82t-81-55t-100-20q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10h256zm384-256q-53 0-99 20t-82 55t-55 81t-20 100h256q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10zm384 384h-640v1024h640V896z"/>`),
		g.Group(children),
	)
}