package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m338 611l494 305l-342 285L0 882zm986 555v108l-490 293v1l-1-1l-1 1v-1l-489-293v-108l147 96l342-284v-2l1 1l1-1v2l343 284zM490 22l342 285l-494 304L0 341zm836 589l338 271l-489 319l-343-285zM1175 22l489 319l-338 270l-494-304z"/>`),
		g.Group(children),
	)
}