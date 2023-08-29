package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VehicleSubwayTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.25 3A3.75 3.75 0 0 1 20 6.75v9a3.751 3.751 0 0 1-2.89 3.651l2.462 1.172a.75.75 0 0 1-.55 1.392l-.095-.038L13.83 19.5h-3.661l-5.097 2.427a.75.75 0 1 1-.645-1.354L6.89 19.4A3.752 3.752 0 0 1 4 15.75v-9A3.75 3.75 0 0 1 7.75 3h8.5ZM8 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm8 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm.25-10.5h-8.5A2.25 2.25 0 0 0 5.5 6.75v5.75h13V6.75a2.25 2.25 0 0 0-2.25-2.25Zm-3 1.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1 0-1.5h2.5Z"/>`),
		g.Group(children),
	)
}