package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendPlaneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.946 9.316c-.522-.175-.526-.456.011-.635L21.043 2.32c.529-.176.832.12.684.638l-5.453 19.086c-.151.529-.456.547-.68.045L12 14l6-8l-8 6l-8.054-2.684Z"/>`),
		g.Group(children),
	)
}