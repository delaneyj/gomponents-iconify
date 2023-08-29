package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurgicalMaskLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.485 3.121l7.758 1.94a1 1 0 0 1 .757.97V7h1c1.1 0 2 .9 2 2v3a3 3 0 0 1-3 3h-.421a5.999 5.999 0 0 1-2.896 3.158l-4.789 2.395a2 2 0 0 1-1.788 0l-4.79-2.395A5.999 5.999 0 0 1 3.422 15H3a3 3 0 0 1-3-3V9a2 2 0 0 1 2-2h1v-.97a1 1 0 0 1 .757-.97l7.758-1.939a2 2 0 0 1 .97 0ZM12 5.061l-7 1.75v5.98a4 4 0 0 0 2.211 3.578L12 18.764l4.789-2.395A4 4 0 0 0 19 12.792v-5.98l-7-1.75ZM3 9H2v3a1 1 0 0 0 1 1V9Zm19 0h-1v4a1 1 0 0 0 1-1V9Z"/>`),
		g.Group(children),
	)
}