package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fingerprint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7a1 1 0 1 1-2 0a7 7 0 1 1 14 0v2a1 1 0 0 1-2 0V7A5 5 0 1 0 2 7zm3 3a1 1 0 0 1-2 0V7a4 4 0 1 1 8 0a1 1 0 0 1-2 0a2 2 0 1 0-4 0v3zm-2 3a1 1 0 0 1 2 0a2 2 0 1 0 4 0v-3a1 1 0 1 1 2 0v3a4 4 0 1 1-8 0zm3-5a1 1 0 1 1 2 0v4a1 1 0 0 1-2 0V8zm-6 3a1 1 0 0 1 2 0v2a5 5 0 0 0 10 0a1 1 0 0 1 2 0a7 7 0 0 1-14 0v-2z"/>`),
		g.Group(children),
	)
}