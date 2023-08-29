package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.37 14.75L20 10V6a4 4 0 0 0-8 0v4l-9.37 4.75a1 1 0 0 0-.63.93v5a1 1 0 0 0 1 1a1 1 0 0 0 .29-.05L12 18v5l-3.45 1.72a1 1 0 0 0-.55.9V29a1 1 0 0 0 1 1a.9.9 0 0 0 .28 0L16 28l6.72 2a.9.9 0 0 0 .28 0a1 1 0 0 0 1-1v-3.38a1 1 0 0 0-.55-.9L20 23v-5l8.71 3.61a1 1 0 0 0 .29.05a1 1 0 0 0 1-1v-5a1 1 0 0 0-.63-.91ZM28 19.15L18 15v9.24l4 2v1.43l-6-1.75l-6 1.75v-1.43l4-2V15L4 19.15V16.3l10-5.07V6a2 2 0 0 1 4 0v5.23l10 5.07Z"/>`),
		g.Group(children),
	)
}