package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraduationCapFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2L0 9l12 7l10-5.833V17.5h2V9L12 2ZM3.999 13.49V18a9.985 9.985 0 0 0 8 4A9.985 9.985 0 0 0 20 18v-4.509l-8 4.667l-8-4.668Z"/>`),
		g.Group(children),
	)
}