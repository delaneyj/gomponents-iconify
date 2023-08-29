package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSignHashtagSharpSignHashtagTag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 4.25h13m-13 5.5h13M11.25.5l-2.5 13m-3-13l-2.5 13"/>`),
		g.Group(children),
	)
}