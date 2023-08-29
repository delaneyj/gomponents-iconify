package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 17.143V24h24v-6.857zm5.143 5.143H1.714v-1.714h3.429zm12-17.143h-3.429v8.571h-3.429V5.143H6.856L11.999 0z"/>`),
		g.Group(children),
	)
}