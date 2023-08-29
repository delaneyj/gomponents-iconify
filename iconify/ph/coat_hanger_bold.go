package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoatHangerBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m244 168l-96-72l19.2-14.4A12 12 0 0 0 172 72a44 44 0 0 0-87.66-5.48a12 12 0 1 0 23.82 3a20 20 0 0 1 39.09-2.92L121 86.24c-.15.1-.29.21-.43.32L12 168a20 20 0 0 0 12 36h208a20 20 0 0 0 12-36ZM36 180l92-69l92 69Z"/>`),
		g.Group(children),
	)
}