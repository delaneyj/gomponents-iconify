package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelRightTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22.246 4a3.75 3.75 0 0 1 3.75 3.75v11.5a3.75 3.75 0 0 1-3.75 3.75H5.75A3.75 3.75 0 0 1 2 19.25V7.75A3.75 3.75 0 0 1 5.75 4h16.496Zm2.25 3.75a2.25 2.25 0 0 0-2.25-2.25H19v16h3.246a2.25 2.25 0 0 0 2.25-2.25V7.75ZM17.5 21.5v-16H5.75A2.25 2.25 0 0 0 3.5 7.75v11.5a2.25 2.25 0 0 0 2.25 2.25H17.5Z"/>`),
		g.Group(children),
	)
}