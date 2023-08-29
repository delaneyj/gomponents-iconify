package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Font(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M335.151 763.202h435.208L552.753 199.336L335.147 763.202M.004 1192.852v-84.182h104.038L526.542 7.148h133.42l423.294 1101.521H1200v84.182H768.761v-84.182h131.834l-99.271-260.488H302.581l-99.272 260.488h130.244v84.182H0"/>`),
		g.Group(children),
	)
}