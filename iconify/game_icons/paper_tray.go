package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperTray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M87.902 40.998L42.25 391.002h138.566l32 48h86.368l32-48h138.568L424.098 40.998h-98.536l-16 32H202.587l-16.967-32H87.902zm-46.904 368v78.004h430.004v-78.004H340.816l-32 48H203.184l-32-48H40.998z"/>`),
		g.Group(children),
	)
}