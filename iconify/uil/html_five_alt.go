package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlFiveAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.468 2.325A1 1 0 0 0 20.73 2H3.27a1 1 0 0 0-.996 1.089l1.52 17a1 1 0 0 0 .728.874l7.2 2a.996.996 0 0 0 .268.037a1.012 1.012 0 0 0 .267-.036l7.22-2a1 1 0 0 0 .729-.875l1.52-17a1 1 0 0 0-.258-.764Zm-3.193 16.896l-6.284 1.741l-6.266-1.74L4.363 4h15.274ZM7.82 13h6.895l-.327 3.271l-2.568.917l-3.164-1.13a1 1 0 1 0-.673 1.884l3.5 1.25a1.003 1.003 0 0 0 .673 0l3.5-1.25a1 1 0 0 0 .659-.842l.5-5a1 1 0 0 0-.995-1.1H8.725l-.3-3h7.895a1 1 0 0 0 0-2h-9a1 1 0 0 0-.995 1.1l.5 5a1 1 0 0 0 .995.9Z"/>`),
		g.Group(children),
	)
}