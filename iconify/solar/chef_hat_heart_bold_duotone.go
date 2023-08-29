package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChefHatHeartBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5.586 21.414c-.502-.502-.574-2.017-.584-3.414h13.997c-.01 1.397-.082 2.912-.584 3.414C17.829 22 16.884 22 15 22H9c-1.885 0-2.828 0-3.414-.586Z"/><path d="M2 10a5 5 0 0 1 5.737-4.946a4.502 4.502 0 0 1 8.526 0A5 5 0 0 1 19 14.584V18H5v-3.416A5.001 5.001 0 0 1 2 10Z" opacity=".5"/><path d="M11.043 13.67C10.165 13.024 9 11.984 9 11c0-1.673 1.65-2.297 3-1.005c1.35-1.292 3-.668 3 1.005c0 .985-1.165 2.025-2.043 2.67c-.42.307-.63.461-.957.461c-.328 0-.537-.154-.957-.462Z"/></g>`),
		g.Group(children),
	)
}