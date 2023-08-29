package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendPlaneLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.726 2.957l-5.453 19.086c-.15.529-.475.553-.717.07L11 13L1.923 9.37c-.51-.205-.504-.51.034-.689L21.043 2.32c.528-.176.832.12.683.638Zm-2.69 2.14L6.811 9.17l5.637 2.255l3.04 6.081l3.546-12.41Z"/>`),
		g.Group(children),
	)
}