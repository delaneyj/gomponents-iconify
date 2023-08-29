package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodLaughSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 10c-1.246 0-2.233-.835-2.454-2h4.908c-.221 1.165-1.208 2-2.454 2Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm.5 1a.5.5 0 0 0-.5.5C4 9.47 5.53 11 7.5 11S11 9.47 11 7.5a.5.5 0 0 0-.5-.5h-6ZM11 6h-1V5h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}