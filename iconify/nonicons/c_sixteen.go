package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.586.102a.75.75 0 0 1 .756 0l6.214 3.625a.75.75 0 0 1 .373.648v7.25a.75.75 0 0 1-.373.648l-6.214 3.625a.75.75 0 0 1-.756 0l-6.214-3.625A.75.75 0 0 1 1 11.625v-7.25a.75.75 0 0 1 .372-.648L7.586.102Zm.378 14.28l5.465-3.188V4.806L7.964 1.618L2.5 4.806v6.388l5.464 3.188Z"/><path fill="currentColor" d="M8 5.5a2.5 2.5 0 1 0 1.81 4.225a.75.75 0 1 1 1.086 1.035a4 4 0 1 1-.01-5.53a.75.75 0 0 1-1.082 1.04A2.49 2.49 0 0 0 8 5.5Z"/>`),
		g.Group(children),
	)
}