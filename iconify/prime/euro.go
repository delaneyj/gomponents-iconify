package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 20.75h-.15a8.75 8.75 0 0 1-5.93-15a8.54 8.54 0 0 1 6.23-2.46a8.75 8.75 0 0 1 6 2.5a.75.75 0 0 1 0 1.06a.75.75 0 0 1-1.06 0a7.26 7.26 0 1 0-.19 10.53l.22-.21a.79.79 0 0 1 1.09 0a.7.7 0 0 1 .05 1l-.05.05l-.29.28A8.72 8.72 0 0 1 13 20.75Z"/><path fill="currentColor" d="M17 11.25H3a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5Zm-1.5 3H3a.75.75 0 0 1 0-1.5h12.5a.75.75 0 0 1 0 1.5Z"/>`),
		g.Group(children),
	)
}