package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KiaAccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.972 9.588a20.423 20.423 0 0 1 10.48 8.444M26.181 4.5a25.043 25.043 0 0 1 16.38 13.167m-22.428-3.555c-8.115 0-14.694 6.579-14.694 14.694S12.018 43.5 20.133 43.5c.834 0 1.65-.075 2.446-.208v-8.304a6.648 6.648 0 1 1 4.202-6.183v14.629l8.046.066V28.806c0-8.115-6.579-14.694-14.694-14.694Z"/>`),
		g.Group(children),
	)
}