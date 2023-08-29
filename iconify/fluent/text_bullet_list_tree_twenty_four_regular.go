package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListTreeTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.25 17.5a1.25 1.25 0 1 1 0 2.499a1.25 1.25 0 0 1 0-2.499Zm3.5.5h10.5a.75.75 0 0 1 .102 1.493l-.102.007h-10.5a.75.75 0 0 1-.102-1.493L10.75 18h10.5h-10.5Zm-7.5-7a1.25 1.25 0 1 1 0 2.499a1.25 1.25 0 0 1 0-2.499Zm3.5.5h14.5a.75.75 0 0 1 .102 1.493L21.25 13H6.75a.75.75 0 0 1-.102-1.493l.102-.007h14.5h-14.5Zm-3.5-7a1.25 1.25 0 1 1 0 2.499a1.25 1.25 0 0 1 0-2.499Zm3.5.5h14.5a.75.75 0 0 1 .102 1.493l-.102.007H6.75a.75.75 0 0 1-.102-1.493L6.75 5h14.5h-14.5Z"/>`),
		g.Group(children),
	)
}