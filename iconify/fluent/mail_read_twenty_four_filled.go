package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReadTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m2 10.127l9.653 5.038a.75.75 0 0 0 .694 0L22 10.128v7.622a3.25 3.25 0 0 1-3.065 3.245L18.75 21H5.25a3.25 3.25 0 0 1-3.245-3.066L2 17.75v-7.623Zm1.1-1.958l8.517-5.064a.75.75 0 0 1 .662-.051l.104.051L20.9 8.17c.235.14.439.319.605.526L12 13.655l-9.505-4.96a2.25 2.25 0 0 1 .435-.414l.17-.112l8.517-5.064L3.1 8.17Z"/>`),
		g.Group(children),
	)
}