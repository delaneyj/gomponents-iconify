package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpellCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 22h2l-5-12h-2l-5 12h2l1.24-3h5.53Zm-5.93-5l1.82-4.42h.25L18 17Z"/><path fill="currentColor" d="M12 28H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v11h-2V6H6v20h6zm11-.82l-2.59-2.59L19 26l4 4l7-7l-1.41-1.41L23 27.18z"/>`),
		g.Group(children),
	)
}