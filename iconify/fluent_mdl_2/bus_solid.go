package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 640v256h-128v1024q0 27-10 50t-27 40t-41 28t-50 10h-128q-27 0-50-10t-40-27t-28-41t-10-50H512q0 27-10 50t-27 40t-41 28t-50 10H256q-27 0-50-10t-40-27t-28-41t-10-50V896H0V640h128V320q0-40 15-75t41-61t61-41t75-15h1280q40 0 75 15t61 41t41 61t15 75v320h128zM256 1152h640V640H256v512zm256 320q0-26-19-45t-45-19q-26 0-45 19t-19 45v64h128v-64zm1152 192H256v128h1408v-128zm-128-128v-64q0-26-19-45t-45-19q-26 0-45 19t-19 45v64h128zm128-896h-640v512h640V640z"/>`),
		g.Group(children),
	)
}