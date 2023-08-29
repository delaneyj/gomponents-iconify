package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.5 6a2.5 2.5 0 0 0-2.446 1.985C7.172 7.86 5.466 6.963 5.081 5.931a2.5 2.5 0 1 0-1.08.019v4.1a2.5 2.5 0 1 0 1 0V7.466a6.985 6.985 0 0 0 4.047 1.52A2.5 2.5 0 1 0 11.5 6ZM3 3.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm3 9a1.5 1.5 0 1 1-2.999 0a1.5 1.5 0 0 1 3 0Zm5.5-2.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}