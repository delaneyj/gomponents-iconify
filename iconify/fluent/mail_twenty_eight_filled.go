package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m3 10.124l10.654 5.541a.75.75 0 0 0 .692 0L25 10.125v9.625a3.25 3.25 0 0 1-3.066 3.245L21.75 23H6.25a3.25 3.25 0 0 1-3.245-3.066L3 19.75v-9.626ZM6.25 5h15.5a3.25 3.25 0 0 1 3.245 3.066L25 8.25v.184l-11 5.72l-11-5.72V8.25a3.25 3.25 0 0 1 3.066-3.245L6.25 5h15.5h-15.5Z"/>`),
		g.Group(children),
	)
}