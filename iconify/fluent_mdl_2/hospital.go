package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 768v256H640v256h256v256h256v-256h256v-256h-256V768H896zm384 128h256v512h-256v256H768v-256H512V896h256V640h512v256zm576-512q40 0 75 15t61 41t41 61t15 75v1152q0 40-15 75t-41 61t-61 41t-75 15H192q-40 0-75-15t-61-41t-41-61t-15-75V576q0-40 15-75t41-61t61-41t75-15h320q0-37-1-82t9-83t37-65t83-26h768q37 0 61 12t38 32t20 46t9 55t1 57t-1 54h320zm-1216 0h768V256H640v128zm1280 192q0-26-19-45t-45-19H192q-26 0-45 19t-19 45v1152q0 26 19 45t45 19h1664q26 0 45-19t19-45V576z"/>`),
		g.Group(children),
	)
}