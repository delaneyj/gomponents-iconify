package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveSeries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M21.56 15.1l-3.48-4.35a2 2 0 0 0-1.56-.75H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V16.35a2 2 0 0 0-.44-1.25zM9 12h6v3H9zm6 16H9v-6h6zm2 0v-6a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v6H4V12h3v3a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-2.4l3 3.75V28z" fill="currentColor"/><path d="M27 21h-2V7H11V5h14a2 2 0 0 1 2 2z" fill="currentColor"/><path d="M32 14h-2V2H18V0h12a2 2 0 0 1 2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}