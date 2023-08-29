package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHandUp0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M41 38H19v6h22v-6Z"/><path d="M19 38C12.481 30.877 8.74 26.75 7.778 25.616c-1.443-1.7-.837-3.62 2.775-3.62s5.695 5.285 8.447 5.285c.017.004.018-6.755.003-20.277A3 3 0 0 1 22.001 4h.003a3.004 3.004 0 0 1 3.005 3.004v8.01c7.972 1.208 12.307 1.875 13.003 2C39.057 17.202 41 18.2 41 21.068V38H19Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHandUp0)"/>`),
		g.Group(children),
	)
}