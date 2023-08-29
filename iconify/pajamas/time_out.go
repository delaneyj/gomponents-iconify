package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.229.199a8 8 0 0 1 9.727 6.964a.75.75 0 0 1-1.492.157a6.5 6.5 0 1 0-7.132 7.146a.75.75 0 1 1-.154 1.492a8 8 0 0 1-.95-15.76ZM8 3a.75.75 0 0 1 .75.75V9h-4a.75.75 0 0 1 0-1.5h2.5V3.75A.75.75 0 0 1 8 3Zm2.22 7.22a.75.75 0 0 1 1.06 0L13 11.94l1.72-1.72a.75.75 0 1 1 1.06 1.06L14.06 13l1.72 1.72a.75.75 0 1 1-1.06 1.06L13 14.06l-1.72 1.72a.75.75 0 1 1-1.06-1.06L11.94 13l-1.72-1.72a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}