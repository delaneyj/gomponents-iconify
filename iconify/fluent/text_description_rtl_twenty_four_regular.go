package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDescriptionRtlTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.25 17H8.75a.75.75 0 0 0-.102 1.493l.102.007h12.5a.75.75 0 0 0 .102-1.493L21.25 17H8.75h12.5Zm0-4H2.75a.75.75 0 0 0-.102 1.493l.102.007h18.5a.75.75 0 0 0 .102-1.493L21.25 13H2.75h18.5Zm0-4H2.75a.75.75 0 0 0-.102 1.493l.102.007h18.5a.75.75 0 0 0 .102-1.493L21.25 9H2.75h18.5Zm0-4H2.75a.75.75 0 0 0-.102 1.493l.102.007h18.5a.75.75 0 0 0 .102-1.493L21.25 5H2.75h18.5Z"/>`),
		g.Group(children),
	)
}