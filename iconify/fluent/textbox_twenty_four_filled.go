package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextboxTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 5.75A2.75 2.75 0 0 0 18.25 3H5.75A2.75 2.75 0 0 0 3 5.75v12.5A2.75 2.75 0 0 0 5.75 21h12.5A2.75 2.75 0 0 0 21 18.25V5.75Zm-6.75 5.75a.75.75 0 0 1 .102 1.493L14.25 13h-7.5a.75.75 0 0 1-.102-1.493l.102-.007h7.5Zm-7.5 4h10.5a.75.75 0 0 1 .102 1.493L17.25 17H6.75a.75.75 0 0 1-.102-1.493l.102-.007Zm10.5-8a.75.75 0 0 1 .102 1.493L17.25 9H6.75a.75.75 0 0 1-.102-1.493L6.75 7.5h10.5Z"/>`),
		g.Group(children),
	)
}