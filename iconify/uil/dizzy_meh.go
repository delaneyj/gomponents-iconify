package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DizzyMeh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 11.71l.29-.3l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29l.3-.29a1 1 0 0 0-1.46-1.42l-.29.3l-.25-.3a1 1 0 1 0-1.46 1.42l.3.29l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0ZM15 14H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm5-11.71a1 1 0 0 0-1.42 0l-.29.3l-.29-.3a1 1 0 0 0-1.42 1.42l.3.29l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l.29-.3l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29l.3-.29a1 1 0 0 0 0-1.42Z"/>`),
		g.Group(children),
	)
}