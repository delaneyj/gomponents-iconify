package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M210.5 419.9h93.1V233.7h139.6L257 1L70.8 233.7h139.6v186.2zm256-46.5v93.1h-419v-93.1H1V513h512V373.4h-46.5z"/>`),
		g.Group(children),
	)
}