package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.562 4.148a7.03 7.03 0 0 1 2.876 0l1.683-1.684l1.414 1.415l-1.05 1.05a7.031 7.031 0 0 1 2.841 3.07H21v2h-2.07c.046.327.07.661.07 1v1h2v2h-2v1c0 .34-.024.674-.07 1H21v2h-2.674a7 7 0 0 1-12.652 0H3v-2h2.07A7.06 7.06 0 0 1 5 15v-1H3v-2h2v-1c0-.339.024-.673.07-1H3V8h2.674a7.03 7.03 0 0 1 2.84-3.07l-1.05-1.05L8.88 2.464l1.683 1.684ZM12 6a5 5 0 0 0-5 5v4a5 5 0 0 0 10 0v-4a5 5 0 0 0-5-5Zm-3 8h6v2H9v-2Zm0-4h6v2H9v-2Z"/>`),
		g.Group(children),
	)
}