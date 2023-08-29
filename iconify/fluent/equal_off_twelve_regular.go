package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualOffTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.854 1.146a.5.5 0 1 0-.708.708L3.293 4H2a.5.5 0 0 0 0 1h2.293l2 2H2a.5.5 0 0 0 0 1h5.293l2.853 2.854a.5.5 0 0 0 .708-.708l-9-9ZM9.12 7l.988.988A.5.5 0 0 0 10 7h-.879Zm-3-3l1 1H10a.5.5 0 0 0 0-1H6.121Z"/>`),
		g.Group(children),
	)
}