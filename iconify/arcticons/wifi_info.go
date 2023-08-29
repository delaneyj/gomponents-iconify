package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.595 18.816c2.538-2.658 6.783-4.26 11.337-4.282s8.827 1.544 11.415 4.178m-22.847.266l11.437 14.488L35.5 18.876m-11.59 3.295v4.642"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="23.91" cy="19.215" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}