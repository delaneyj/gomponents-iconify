package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighPriority(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.576 1.424a.6.6 0 0 1 .848 0l10.152 10.152a.6.6 0 0 1 0 .848L12.424 22.576a.6.6 0 0 1-.848 0L1.424 12.424a.6.6 0 0 1 0-.848L11.576 1.424ZM12 8v4m0 4.01l.01-.011"/>`),
		g.Group(children),
	)
}