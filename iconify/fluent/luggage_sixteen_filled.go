package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.834 1a.5.5 0 0 0 0 1H6v1H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2v.5a.5.5 0 0 0 1 0V14h4v.5a.5.5 0 0 0 1 0V14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-1V2h.167a.5.5 0 0 0 0-1H5.834ZM7 3V2h2v1H7ZM5 6.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}