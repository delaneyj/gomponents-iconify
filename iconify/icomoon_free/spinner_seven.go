package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinnerSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.5 14.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 6.5 14.5zM0 8a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 0 8zm13 0a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 13 8zM1.904 3.404a1.5 1.5 0 1 1 3.001.001a1.5 1.5 0 0 1-3.001-.001zm9.192 9.192a1.5 1.5 0 1 1 3.001.001a1.5 1.5 0 0 1-3.001-.001zm-9.192 0a1.5 1.5 0 1 1 3.001.001a1.5 1.5 0 0 1-3.001-.001zm9.192-9.192a1.5 1.5 0 1 1 3.001.001a1.5 1.5 0 0 1-3.001-.001z"/>`),
		g.Group(children),
	)
}