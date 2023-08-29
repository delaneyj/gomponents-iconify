package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwinColSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 2v12h16V2H0zm7 11H1V3h6v10zm8 0H9V3h6v10z"/><path fill="currentColor" d="M10 4h4v1h-4V4zM2 4h4v1H2V4zm0 2h4v1H2V6zm0 2h4v1H2V8z"/>`),
		g.Group(children),
	)
}