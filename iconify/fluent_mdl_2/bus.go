package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 640v256h-128v960q0 34-3 68t-16 62t-38 45t-71 17q-25 0-53 1t-57-1t-55-8t-47-21t-32-38t-12-61H512q0 27-10 50t-27 40t-41 28t-50 10q-69 0-118-1t-79-17t-45-56t-14-118V896H0V640h128V320q0-40 15-75t41-61t61-41t75-15h1280q40 0 75 15t61 41t41 61t15 75v320h128zm-256 640H256v256h128v-40q0-22 4-42t18-33t42-13q28 0 41 13t19 32t5 42t-1 41h896v-40q0-22 4-42t18-33t42-13q28 0 41 13t19 32t5 42t-1 41h128v-256zm-768-128V640H256v512h640zm768-512h-640v512h640V640zM320 256q-26 0-45 19t-19 45v192h1408V320q0-26-19-45t-45-19H320zm1344 1536v-128H256v128h1408z"/>`),
		g.Group(children),
	)
}