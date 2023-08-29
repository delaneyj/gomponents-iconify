package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleWeather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.964 40.536L24 44.5l-6.004-6.004H9.504v-8.492h0L3.5 24l6.004-6.004V9.504h8.492L24 3.5l6.004 6.004h8.492v8.492L44.5 24l-3.011 3.011"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.808 36.764a7.727 7.727 0 1 1 6.706 3.887l-12.038-.163a4.636 4.636 0 0 1 0-9.272l4.506-.014m-8.918-11.615h6.216a6.248 6.248 0 0 1-6.096 6.397l-.12.002a6.4 6.4 0 0 1 0-12.799a6.303 6.303 0 0 1 3.173.842"/>`),
		g.Group(children),
	)
}