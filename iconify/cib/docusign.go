package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Docusign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.688 4.411h6.625v8.828h4.411L16 22.067l-7.724-8.828h4.411zM0 25.38h32v2.203H0z"/>`),
		g.Group(children),
	)
}