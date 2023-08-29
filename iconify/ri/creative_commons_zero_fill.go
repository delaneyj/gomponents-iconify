package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsZeroFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.52 0 10 4.48 10 10s-4.48 10-10 10S2 17.52 2 12S6.48 2 12 2Zm0 4c-2.761 0-5 2.686-5 6s2.239 6 5 6s5-2.686 5-6s-2.239-6-5-6Zm2.325 3.472c.422.69.675 1.57.675 2.528c0 2.21-1.343 4-3 4c-.441 0-.86-.127-1.237-.355l3.562-6.173ZM12 8c.441 0 .86.127 1.237.355l-3.562 6.173C9.253 13.838 9 12.958 9 12c0-2.21 1.343-4 3-4Z"/>`),
		g.Group(children),
	)
}