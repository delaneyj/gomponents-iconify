package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.49 3L10.74.37A1.22 1.22 0 0 0 9.86 0h-4a1.25 1.25 0 0 0-1.22 1.25v11a1.25 1.25 0 0 0 1.25 1.25h6.72a1.25 1.25 0 0 0 1.25-1.25V3.88a1.22 1.22 0 0 0-.37-.88zm-.88 9.25H5.89v-11h2.72v2.63a1.25 1.25 0 0 0 1.25 1.25h2.75zm0-8.37H9.86V1.25l2.75 2.63z"/><path fill="currentColor" d="M10.11 14.75H3.39v-11H4V2.5h-.61a1.25 1.25 0 0 0-1.25 1.25v11A1.25 1.25 0 0 0 3.39 16h6.72a1.25 1.25 0 0 0 1.25-1.25v-.63h-1.25z"/>`),
		g.Group(children),
	)
}