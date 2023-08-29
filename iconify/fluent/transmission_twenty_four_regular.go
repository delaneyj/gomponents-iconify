package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransmissionTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2a2.75 2.75 0 0 0-.75 5.396v3.854a.75.75 0 0 0 1.5 0V7.396A2.751 2.751 0 0 0 12 2Zm-1.25 2.75a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0ZM3 6.5a2.5 2.5 0 0 1 5 0V10h2.25v1.5H6.5v-5a1 1 0 0 0-2 0v12a1 1 0 1 0 2 0v-5H11v5a1 1 0 1 0 2 0v-5h6a.5.5 0 0 0 .5-.5V6.5a1 1 0 1 0-2 0v5h-3.75V10H16V6.5a2.5 2.5 0 0 1 5 0V13a2 2 0 0 1-2 2h-4.5v3.5a2.5 2.5 0 0 1-5 0V15H8v3.5a2.5 2.5 0 0 1-5 0v-12Z"/>`),
		g.Group(children),
	)
}