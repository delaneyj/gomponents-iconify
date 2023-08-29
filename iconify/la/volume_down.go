package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 4.594L13.281 6.28L8.562 11H4v10h4.563l4.718 4.719L15 27.406zm-2 4.843v13.126L9.719 19.28L9.406 19H6v-6h3.406l.313-.281zm5.5 2.594l-1.438 1.438c.579.695.938 1.558.938 2.531c0 .973-.36 1.836-.938 2.531L18.5 19.97A5.956 5.956 0 0 0 20 16a5.96 5.96 0 0 0-1.5-3.969z"/>`),
		g.Group(children),
	)
}