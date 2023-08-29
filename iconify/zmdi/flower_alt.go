package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 453q0-79 56-135.5T384 261q0 80-56 136t-136 56zM55 203q0-34 31-48q-31-15-31-48q0-22 16-38t38-16q17 0 30 10v-4q0-22 15.5-38T192 5t37.5 16T245 59v4q14-10 30-10q22 0 38 16t16 38q0 33-31 48q31 14 31 48q0 22-16 37.5T275 256q-17 0-30-9v4q0 22-15.5 37.5T192 304t-37.5-15.5T139 251v-4q-14 9-30 9q-22 0-38-15.5T55 203zm137-102q-22 0-37.5 16T139 155t15.5 37.5T192 208t37.5-15.5T245 155t-15.5-38t-37.5-16zM0 261q80 0 136 56.5T192 453q-80 0-136-56T0 261z"/>`),
		g.Group(children),
	)
}