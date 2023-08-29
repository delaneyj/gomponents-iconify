package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Annotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-2.172a1 1 0 0 0-.707.293l-2 2a3 3 0 0 1-4.242 0l-2-2A1 1 0 0 0 7.172 18H5a3 3 0 0 1-3-3V6zm5 1a1 1 0 0 0 0 2h10a1 1 0 1 0 0-2H7zm0 4a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2H7z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}