package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoGitlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.78 1h1.936l2.5 7.333h5.568L17.284 1h1.937l4.454 13.362L12 23.257L.325 14.362L4.78 1Zm.978 3.389l-3.083 9.249L12 20.743l9.325-7.105l-3.083-9.25l-2.026 5.945H7.784L5.758 4.39Z"/>`),
		g.Group(children),
	)
}