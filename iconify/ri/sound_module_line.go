package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundModuleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18v3h-2v-3h-2v-2h6v2h-2ZM5 18v3H3v-3H1v-2h6v2H5Zm6-12V3h2v3h2v2H9V6h2Zm0 4h2v11h-2V10Zm-8 4V3h2v11H3Zm16 0V3h2v11h-2Z"/>`),
		g.Group(children),
	)
}