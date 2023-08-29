package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feEye0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feEye1" fill="currentColor"><path id="feEye2" d="M2 12c.945-4.564 5.063-8 10-8s9.055 3.436 10 8c-.945 4.564-5.063 8-10 8s-9.055-3.436-10-8Zm10 5a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0-2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></g>`),
		g.Group(children),
	)
}