package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StormWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M9.76 24.6V7.45h1.13V24.6H9.76zm1.94-10.55v-6.6h8.55v6.6H11.7zm2.36-2h3.81v-2.5h-3.81v2.5z"/>`),
		g.Group(children),
	)
}