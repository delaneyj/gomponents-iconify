package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usagedirect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 36.5v-25m6 25V16.937M24 36.5V13.928M30 36.5V22.267M36 36.5V16.73M3.199 18.545s12.285-1.676 18.408 6.135c6.721 8.575 11.802 3.652 17.17 4.154c4.212.394 5.765 1.532 5.765 1.532"/>`),
		g.Group(children),
	)
}