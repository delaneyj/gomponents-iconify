package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gymroutines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.065 34.474h25.87m0 8.026V26.447M11.065 42.5V26.447m30.242 12.099V30.4M6.693 38.546V30.4m10.6-13.725l6.313 6.314L41.095 5.5m.394 17.488a17.41 17.41 0 0 0-2.167-8.438m-6.894-6.89A17.494 17.494 0 0 0 6.512 22.989"/>`),
		g.Group(children),
	)
}