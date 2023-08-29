package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionRotateNinetyAcTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M4.293 17.293a1 1 0 0 0 0 1.414l2 2a1 1 0 0 0 1.414 0l2-2a1 1 0 0 0-1.414-1.414L8 17.586V4a1 1 0 0 0-2 0v13.586l-.293-.293a1 1 0 0 0-1.414 0zm9 1.414a1 1 0 0 1 1.414-1.414l.293.293V16a1 1 0 1 1 2 0v1.586l.293-.293a1 1 0 0 1 1.414 1.414l-2 2a1 1 0 0 1-1.414 0l-2-2zM20 12a1 1 0 1 0 0-2h-9.25a1 1 0 1 0 0 2H15v1a1 1 0 1 0 2 0v-1h3zm0-8a1 1 0 1 0-2 0v3h-.25c-.895 0-1.87-.184-2.586-.642C14.511 5.939 14 5.255 14 4a1 1 0 1 0-2 0c0 1.945.864 3.26 2.086 4.042c1.158.742 2.56.958 3.664.958H19a1 1 0 0 0 1-1V4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}