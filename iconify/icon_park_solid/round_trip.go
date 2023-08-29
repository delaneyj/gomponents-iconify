package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundTrip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRoundTrip0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="m9 24l-4-4s-1.4 4.728-1.076 7.578C4.248 30.428 7.274 32.574 10 31c2.727-1.574 34-21 34-21l-9-2L9 24Z"/><path d="m26 13l-15.202-1.615L8 13l7 7m14 24l-4-5h17v-4m-10-7l4 5H19v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRoundTrip0)"/>`),
		g.Group(children),
	)
}