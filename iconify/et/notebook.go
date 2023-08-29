package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M23.5 32c.869 0 1.5-.631 1.5-1.5v-28c0-.869-.631-1.5-1.5-1.5h-20C2.631 1 2 1.631 2 2.5v1.167a.5.5 0 0 0 1 0V2.5c0-.313.187-.5.5-.5h4.562v29H3.5c-.313 0-.5-.187-.5-.5v-.917a.5.5 0 0 0-1 0v.917c0 .869.631 1.5 1.5 1.5h20zM9 2h14.5c.313 0 .5.187.5.5v28c0 .313-.187.5-.5.5H9V2zm18 27.5v-26a.5.5 0 0 0-1 0v26a.5.5 0 0 0 1 0z"/><path d="M.5 11h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1zm0-4h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1zm0 8h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1zm0 8h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1zm0-4h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1zm0 8h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1z"/></g>`),
		g.Group(children),
	)
}