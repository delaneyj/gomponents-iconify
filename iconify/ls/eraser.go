package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eraser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M68 276L322 22c23-23 58-23 81 0l246 246c23 23 23 59 0 82L395 604c-90 90-236 90-327 0c-90-90-90-238 0-328zm286 287L109 317c-68 68-68 178 0 246s177 68 245 0z"/>`),
		g.Group(children),
	)
}