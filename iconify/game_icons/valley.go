package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Valley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M253.844 20.72L189.03 192.312l-49.31-47.188l-60.5 150.844h362.59l-57-94.564l-66.03 68.125l-13.407-13.03l37.938-39.125l-89.47-196.656zM20.47 314.655v178.72h175.75l49.936-78.626l-36.062-34.844l43.875-65.25H20.47zm261.186 0l-35.5 55.313l58.47 47.25l-2.126 76.155h193v-178.72H281.656z"/>`),
		g.Group(children),
	)
}