package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.906 4.781L4.687 16l11.22 11.219l1.405-1.438L7.533 16l9.78-9.781zm7 0L11.688 16l11.218 11.219l1.407-1.438L14.53 16l9.781-9.781z"/>`),
		g.Group(children),
	)
}