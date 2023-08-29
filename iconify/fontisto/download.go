package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 17.143V24h24v-6.857zm5.143 5.143H1.714v-1.714h3.429zM6.857 8.571h3.429V0h3.429v8.571h3.429l-5.143 5.143z"/>`),
		g.Group(children),
	)
}