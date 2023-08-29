package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oroaming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.897 33.282v-5.075l-14.625-9.834v-10.6a3.272 3.272 0 0 0-6.544 0v10.6L6.103 28.207v5.075l14.625-5.168v8.703l-4.875 4.212V43.5h16.294v-2.47l-4.875-4.213v-8.703Z"/>`),
		g.Group(children),
	)
}