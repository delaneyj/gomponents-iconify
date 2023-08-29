package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 0v3h3z"/><path fill="currentColor" d="M9 4H5V0H0v12h9zm4 0v3h3z"/><path fill="currentColor" d="M12 4h-2v9H7v3h9V8h-4z"/>`),
		g.Group(children),
	)
}