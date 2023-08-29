package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Script(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18.83 26l2.58-2.58L20 22l-4 4l4 4l1.42-1.41L18.83 26zm8.34 0l-2.58 2.58L26 30l4-4l-4-4l-1.42 1.41L27.17 26z"/><path fill="currentColor" d="M14 28H8V4h8v6a2.006 2.006 0 0 0 2 2h6v6h2v-8a.91.91 0 0 0-.3-.7l-7-7A.909.909 0 0 0 18 2H8a2.006 2.006 0 0 0-2 2v24a2.006 2.006 0 0 0 2 2h6Zm4-23.6l5.6 5.6H18Z"/>`),
		g.Group(children),
	)
}