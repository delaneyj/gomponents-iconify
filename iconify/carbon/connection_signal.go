package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectionSignal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 12h2v18h-2zm-3.67 6.22a7 7 0 0 1 0-10.44l1.34 1.49a5 5 0 0 0 0 7.46zm9.34 0l-1.34-1.49a5 5 0 0 0 0-7.46l1.34-1.49a7 7 0 0 1 0 10.44z"/><path fill="currentColor" d="M8.4 21.8a11 11 0 0 1 0-17.6l1.2 1.6a9 9 0 0 0 0 14.4zm15.2 0l-1.2-1.6a9 9 0 0 0 0-14.4l1.2-1.6a11 11 0 0 1 0 17.6z"/>`),
		g.Group(children),
	)
}