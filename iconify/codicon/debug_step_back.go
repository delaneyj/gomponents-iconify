package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugStepBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.75 5.75v-4h1.5v2.542c1.145-1.359 2.911-2.209 4.84-2.209c3.177 0 5.92 2.307 6.16 5.398l.02.269h-1.5l-.022-.226c-.212-2.195-2.202-3.94-4.656-3.94c-1.736 0-3.244.875-4.05 2.166h2.83v1.5H2.707l-.961-.975V5.75h.003zM8 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}