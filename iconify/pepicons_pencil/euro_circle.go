package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.989 6C11.668 6 9 8.423 9 12c0 3.626 2.716 6.5 5.989 6.5c.817 0 1.595-.177 2.305-.499a.5.5 0 0 1 .412.91a6.562 6.562 0 0 1-2.717.589C11.094 19.5 8 16.106 8 12c0-4.155 3.142-7 6.989-7c1.124 0 2.219.355 3.293 1.087a.5.5 0 1 1-.564.826c-.93-.633-1.83-.913-2.73-.913Z"/><path d="M6 10.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5ZM6 14a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8A.5.5 0 0 1 6 14Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}