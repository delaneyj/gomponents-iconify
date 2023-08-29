package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M39 19.434h-5.384v25.133H39c2.969 0 5.386-2.322 5.386-5.174V24.609c0-2.853-2.417-5.175-5.386-5.175"/><path fill="currentColor" d="M52 2H12C6.477 2 2 6.477 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10M23 49h-4V15h4v34m26-9.607a9.251 9.251 0 0 1-.787 3.738a9.592 9.592 0 0 1-2.143 3.055a10.032 10.032 0 0 1-3.178 2.059A10.302 10.302 0 0 1 39 49H29V15h10c1.348 0 2.657.254 3.893.754c1.19.484 2.26 1.176 3.178 2.059s1.638 1.912 2.143 3.053A9.294 9.294 0 0 1 49 24.609v14.784"/>`),
		g.Group(children),
	)
}