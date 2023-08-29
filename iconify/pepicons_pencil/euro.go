package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.989 3C8.668 3 6 5.423 6 9c0 3.626 2.716 6.5 5.989 6.5c.817 0 1.595-.177 2.305-.499a.5.5 0 0 1 .412.91a6.562 6.562 0 0 1-2.717.589C8.094 16.5 5 13.106 5 9c0-4.155 3.142-7 6.989-7c1.124 0 2.219.355 3.293 1.087a.5.5 0 1 1-.564.826c-.93-.633-1.83-.913-2.73-.913Z"/><path d="M3 7.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5ZM3 11a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8A.5.5 0 0 1 3 11Z"/></g>`),
		g.Group(children),
	)
}