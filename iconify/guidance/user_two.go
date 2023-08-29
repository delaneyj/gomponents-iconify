package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M21.435 16.677V23.5H2.565v-6.823S6.192 14.5 12 14.5c5.806 0 9.435 2.177 9.435 2.177ZM6.92 5.58A5.084 5.084 0 0 1 12.005.5a5.073 5.073 0 0 1 5.075 5.08c0 3.89-4.974 6.533-4.974 6.533h-.214S6.919 9.47 6.919 5.58Z"/>`),
		g.Group(children),
	)
}