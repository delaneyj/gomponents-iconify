package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M465 512H47c-25.957 0-47-21.043-47-47V47C0 21.043 21.043 0 47 0h418c25.957 0 47 21.043 47 47v418c0 25.957-21.043 47-47 47zm-20.584-273.41H91.136V85.504h153.088v118.784h23.552V61.439H67.584v200.704h353.28V415.23H267.776v-99.84h-23.552v123.392h200.192V238.59z"/>`),
		g.Group(children),
	)
}