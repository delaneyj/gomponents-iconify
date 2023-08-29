package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppTitleTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.75 20.5h14.5a.75.75 0 0 1 .102 1.493L19.25 22H4.75a.75.75 0 0 1-.102-1.493l.102-.007h14.5h-14.5ZM16.25 3A3.75 3.75 0 0 1 20 6.75v8.5A3.75 3.75 0 0 1 16.25 19h-8.5A3.75 3.75 0 0 1 4 15.25v-8.5A3.75 3.75 0 0 1 7.75 3h8.5Zm0 1.5h-8.5A2.25 2.25 0 0 0 5.5 6.75v8.5a2.25 2.25 0 0 0 2.25 2.25h8.5a2.25 2.25 0 0 0 2.25-2.25v-8.5a2.25 2.25 0 0 0-2.25-2.25Z"/>`),
		g.Group(children),
	)
}