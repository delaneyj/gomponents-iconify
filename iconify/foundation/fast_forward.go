package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M86.346 49.976c0-.702-.422-1.303-1.026-1.57l.017-.03l-37.39-21.588a1.715 1.715 0 0 0-2.92 1.172h-.005v15.249L16.58 26.788a1.715 1.715 0 0 0-2.92 1.172h-.005v44.031a1.718 1.718 0 0 0 2.948 1.197l28.42-16.408v15.212a1.718 1.718 0 0 0 2.948 1.197l37.545-21.677l-.031-.054c.512-.298.861-.847.861-1.482z"/>`),
		g.Group(children),
	)
}