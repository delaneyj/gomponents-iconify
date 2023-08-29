package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilEuroCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.989 6C11.668 6 9 8.423 9 12c0 3.626 2.716 6.5 5.989 6.5c.817 0 1.595-.177 2.305-.499a.5.5 0 0 1 .412.91a6.562 6.562 0 0 1-2.717.589C11.094 19.5 8 16.106 8 12c0-4.155 3.142-7 6.989-7c1.124 0 2.219.355 3.293 1.087a.5.5 0 1 1-.564.826c-.93-.633-1.83-.913-2.73-.913Z"/><path d="M6 10.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5ZM6 14a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 0 1h-8A.5.5 0 0 1 6 14Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilEuroCircleFilled0)"/></g>`),
		g.Group(children),
	)
}