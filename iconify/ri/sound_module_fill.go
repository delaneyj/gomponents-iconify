package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundModuleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18v3h-2v-3h-2v-3h6v3h-2ZM5 18v3H3v-3H1v-3h6v3H5Zm6-12V3h2v3h2v3H9V6h2Zm0 5h2v10h-2V11Zm-8 2V3h2v10H3Zm16 0V3h2v10h-2Z"/>`),
		g.Group(children),
	)
}