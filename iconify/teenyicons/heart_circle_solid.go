package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm4.146-2.354a1.914 1.914 0 0 1 2.707 0l.647.647l.646-.647a1.914 1.914 0 0 1 2.707 2.708L7.5 11.207L4.146 7.854a1.914 1.914 0 0 1 0-2.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}