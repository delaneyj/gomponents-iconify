package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NSixtyFourOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4.5 4v3M3 5.5h3m3 0h1m1 1h1m-7.5-4h-1a3 3 0 0 0-3 3v4.838a1.162 1.162 0 0 0 2.265.367L3.5 8.5h1.095a1 1 0 0 1 .995.9l.26 2.607a1.657 1.657 0 0 0 3.3 0L9.41 9.4a1 1 0 0 1 .995-.9H11.5l.735 2.205a1.162 1.162 0 0 0 2.265-.367V5.5a3 3 0 0 0-3-3h-1l-1-1h-4l-1 1Z"/>`),
		g.Group(children),
	)
}