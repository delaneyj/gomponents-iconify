package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReadSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.237 1.56a.5.5 0 0 0-.474 0L1.789 4.777a1.499 1.499 0 0 0-.57.54L8 8.933l6.78-3.616a1.498 1.498 0 0 0-.569-.54L8.237 1.56ZM15 6.333L8.235 9.941a.5.5 0 0 1-.47 0L1 6.333V12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6.333Z"/>`),
		g.Group(children),
	)
}