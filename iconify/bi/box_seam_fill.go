package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxSeamFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.528 2.973a.75.75 0 0 1 .472.696v8.662a.75.75 0 0 1-.472.696l-7.25 2.9a.75.75 0 0 1-.557 0l-7.25-2.9A.75.75 0 0 1 0 12.331V3.669a.75.75 0 0 1 .471-.696L7.443.184l.01-.003l.268-.108a.75.75 0 0 1 .558 0l.269.108l.01.003l6.97 2.789ZM10.404 2L4.25 4.461L1.846 3.5L1 3.839v.4l6.5 2.6v7.922l.5.2l.5-.2V6.84l6.5-2.6v-.4l-.846-.339L8 5.961L5.596 5l6.154-2.461L10.404 2Z"/>`),
		g.Group(children),
	)
}