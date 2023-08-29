package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.764 11.284a1.333 1.333 0 0 0-.31-.339A25.511 25.511 0 0 0 5.38 7.082l-.466-.165c-.87-.31-1.79.279-1.907 1.177a30.314 30.314 0 0 0 0 7.812c.118.898 1.037 1.486 1.907 1.178l.466-.166a25.51 25.51 0 0 0 7.073-3.863a1.34 1.34 0 0 0 .31-.339a29.97 29.97 0 0 0 .244 3.19c.118.898 1.037 1.486 1.907 1.178l.466-.166a25.511 25.511 0 0 0 7.073-3.863c.69-.534.69-1.576 0-2.11a25.512 25.512 0 0 0-7.073-3.863l-.466-.165c-.87-.31-1.79.279-1.907 1.177a29.969 29.969 0 0 0-.244 3.19Z"/>`),
		g.Group(children),
	)
}