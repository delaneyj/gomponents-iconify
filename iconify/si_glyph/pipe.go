package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.274 4.644C9.75 6.47 8.318 7.874 6.03 8.494l-.729-2.08c-.43-1.237-5.646.584-5.213 1.821l1.133 3.242c.432 1.236 2.059 1.814 3.543 1.376c1.175-.347 4.087-1.353 7.528-6.349c1.125-1.635 2.566-2.66 3.698-3.006l-.515-1.473s-2.278.316-4.2 2.62z"/>`),
		g.Group(children),
	)
}