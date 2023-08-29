package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.104 21.468v6.519M9.5 29.649v-6.512l3.26 5.953l3.26-7.076v6.509m18.29-3.893c.4.452.9.558 1.598.437l.966-.168c.898-.156 1.626-1.01 1.626-1.909v-.006c0-.899-.728-1.5-1.626-1.344l-1.065.185c-.9.156-1.628-.446-1.628-1.346h0c0-.9.73-1.758 1.631-1.915l.96-.167c.698-.12 1.2-.014 1.599.437m-9.99 7.541l3.26-.566m-3.26-5.954l3.26-.566m-3.26 3.826l2.125-.369m-2.125-2.891v6.52m-6.078-5.463v6.519l3.26-.566"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}