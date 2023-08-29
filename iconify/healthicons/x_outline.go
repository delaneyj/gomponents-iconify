package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M31.695 10.08a3 3 0 0 1 4.61 3.84L27.905 24l8.4 10.08a3 3 0 1 1-4.61 3.84L24 28.687l-7.695 9.235a3 3 0 1 1-4.61-3.841l8.4-10.08l-8.4-10.08a3 3 0 0 1 4.61-3.84L24 19.313l7.695-9.235Zm2.945 1.152a1 1 0 0 0-1.408.128l-.768-.64l.768.64l-8.464 10.156a1 1 0 0 1-1.536 0L14.768 11.36a1 1 0 1 0-1.536 1.28l8.933 10.72a1 1 0 0 1 0 1.28l-8.933 10.72a1 1 0 1 0 1.536 1.28l8.464-10.156a1 1 0 0 1 1.536 0l8.464 10.156a1 1 0 0 0 1.536-1.28l-8.933-10.72a1 1 0 0 1 0-1.28l8.933-10.72a1 1 0 0 0-.128-1.408Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}