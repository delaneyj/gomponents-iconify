package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KotlinSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 2A.75.75 0 0 1 2 1.25h12a.75.75 0 0 1 .53 1.28L9.06 8l5.47 5.47a.75.75 0 0 1-.53 1.28H2a.75.75 0 0 1-.75-.75V2Zm1.5.75v10.5h9.44L7.47 8.53a.75.75 0 0 1 0-1.06l4.72-4.72H2.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}