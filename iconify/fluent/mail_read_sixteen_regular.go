package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReadSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.237 1.56a.5.5 0 0 0-.474 0L1.789 4.777A1.5 1.5 0 0 0 1 6.097V12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6.097a1.5 1.5 0 0 0-.789-1.32L8.237 1.56ZM2.263 5.657L8 2.567l5.737 3.09a.5.5 0 0 1 .157.133L8 8.933L2.106 5.79a.5.5 0 0 1 .157-.133ZM2 6.867L7.765 9.94a.5.5 0 0 0 .47 0L14 6.867V12a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V6.867Z"/>`),
		g.Group(children),
	)
}