package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 6v20h9.563l2.718 2.72l.72.686l.72-.687L19.437 26H29V6H3zm2 2h22v16h-8.406l-.313.28L16 26.563l-2.28-2.28l-.314-.282H5V8zm4 3v2h14v-2H9zm0 4v2h14v-2H9zm0 4v2h10v-2H9z"/>`),
		g.Group(children),
	)
}