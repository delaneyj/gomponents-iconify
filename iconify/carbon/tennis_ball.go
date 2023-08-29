package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TennisBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 4a8.981 8.981 0 0 0-8.169 12.755L2 25.585L3.414 27l8.487-8.487a9.04 9.04 0 0 0 1.586 1.586L10 23.586L11.414 25l3.831-3.831A8.996 8.996 0 1 0 19 4Zm6.906 7.906a7.005 7.005 0 0 1-5.812-5.812a7.005 7.005 0 0 1 5.812 5.812Zm-13.812 2.188a7.005 7.005 0 0 1 5.812 5.812a7.005 7.005 0 0 1-5.812-5.812Zm7.836 5.837a9.01 9.01 0 0 0-7.861-7.862a7.004 7.004 0 0 1 6-6a9.01 9.01 0 0 0 7.861 7.862a7.004 7.004 0 0 1-6 6Z"/>`),
		g.Group(children),
	)
}