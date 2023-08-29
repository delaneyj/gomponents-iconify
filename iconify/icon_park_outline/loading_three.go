package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Zm0-32v3m8.485.515l-2.121 2.121M36 24h-3m-.515 8.485l-2.121-2.121M24 36v-3m-8.485-.515l2.121-2.121M12 24h3m.515-8.485l2.121 2.121"/>`),
		g.Group(children),
	)
}