package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTablet0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTablet1" fill="currentColor" fill-rule="nonzero"><path id="feTablet2" d="M6 4v15h12V4H6Zm0-2h12a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Z"/></g></g>`),
		g.Group(children),
	)
}