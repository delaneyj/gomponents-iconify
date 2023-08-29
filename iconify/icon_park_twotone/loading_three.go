package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLoadingThree0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path d="M24 12v3m8.485.515l-2.121 2.121M36 24h-3m-.515 8.485l-2.121-2.121M24 36v-3m-8.485-.515l2.121-2.121M12 24h3m.515-8.485l2.121 2.121"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLoadingThree0)"/>`),
		g.Group(children),
	)
}