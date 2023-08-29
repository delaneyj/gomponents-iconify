package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeHide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m17.713 13.471l1.863-6.953L17.645 6l-1.565 5.838l1.633 1.633zm6.494 6.494l1.414 1.414L31 16l-7-7l-1.414 1.414L28.172 16l-3.965 3.965zM30 28.586L3.414 2L2 3.414l5.793 5.793L1 16l7 7l1.414-1.414L3.828 16l5.379-5.379l5.677 5.677l-2.461 9.184l1.932.518l2.162-8.069L28.586 30L30 28.586z"/>`),
		g.Group(children),
	)
}