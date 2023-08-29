package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M334.5 0v128h128L334.5 0zm-21.3 0H78.5v512h384V149.3H313.2V0zm-21.4 448h-42.7v-42.7h42.7V448zm0-64h-42.7l-21.3-149.3h85.3L291.8 384z"/>`),
		g.Group(children),
	)
}