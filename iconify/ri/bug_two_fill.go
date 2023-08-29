package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.07 16A7.06 7.06 0 0 1 5 15v-1H3v-2h2v-1c0-.34.024-.674.07-1H3V8h2.674a7.03 7.03 0 0 1 2.84-3.072l-1.05-1.05L8.88 2.465l1.683 1.684a7.03 7.03 0 0 1 2.876 0l1.683-1.684l1.414 1.415l-1.05 1.05A7.031 7.031 0 0 1 18.327 8H21v2h-2.07c.046.327.07.661.07 1v1h2v2h-2v1c0 .34-.024.674-.07 1H21v2h-2.674a7 7 0 0 1-12.652 0H3v-2h2.07ZM9 10v2h6v-2H9Zm0 4v2h6v-2H9Z"/>`),
		g.Group(children),
	)
}