package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.67 1.658c2.103-.494 4.165.222 5.135 1.434c.896 1.119 1.303 2.282 1.17 3.998c-.095 1.237-.5 2.154-1.298 3.163c-.147.186-.308.377-.511.609l-.548.618c-.8.903-3.155 3.163-7.103 6.82a.764.764 0 0 1-1.036-.004C6.09 15.12 3.896 13.028 2.89 12.014C1.033 10.14.326 9.05.07 7.293c-.27-1.84.245-3.349 1.398-4.521c1.082-1.1 3.017-1.508 4.928-1.12c1.119.228 2.31.96 3.6 2.181c1.101-1.129 2.327-1.86 3.675-2.175Z"/>`),
		g.Group(children),
	)
}