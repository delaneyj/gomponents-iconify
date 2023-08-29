package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MixedRainHail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="24.5" cy="25.5" r="1.5" fill="currentColor"/><circle cx="21.5" cy="29.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="m15.868 30.496l-1.736-.992L17.276 24h-5.999l4.855-8.496l1.736.992L14.723 22h6.001l-4.856 8.496z"/><path fill="currentColor" d="M9 32a1 1 0 0 1-.894-1.447l2-4a1 1 0 1 1 1.788.894l-2 4A.998.998 0 0 1 9 32zM24.8 9.136a8.994 8.994 0 0 0-17.6 0a6.493 6.493 0 0 0 .23 12.768l-1.324 2.649a1 1 0 1 0 1.788.894l2-4a1 1 0 0 0-.446-1.341A.979.979 0 0 0 9 20.01V20h-.5a4.497 4.497 0 0 1-.356-8.981l.816-.064l.099-.812a6.994 6.994 0 0 1 13.883 0l.099.812l.815.064A4.497 4.497 0 0 1 23.5 20H23v2h.5a6.497 6.497 0 0 0 1.3-12.864z"/>`),
		g.Group(children),
	)
}