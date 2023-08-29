package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodSadSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm6 0h1V5h-1v1Zm-5.1 4.3a3.25 3.25 0 0 1 5.2 0l.8-.6c-1.7-2.267-5.1-2.267-6.8 0l.8.6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}