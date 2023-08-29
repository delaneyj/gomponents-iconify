package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreeButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M52 2H12C6.477 2 2 6.477 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10zM18 26h-5.09v4.5H18v3h-5.09V41H10V23h8v3zm12.475 15h-3.021l-2.471-7.5h-1.125V41H21V23h5c2.758 0 5 2.355 5 5.25c0 2.197-1.293 4.084-3.121 4.865L30.475 41zM42 26h-5.09v4.5H42v3h-5.09V38H42v3h-8V23h8v3zm12 0h-5.09v4.5H54v3h-5.09V38H54v3h-8V23h8v3z"/><path fill="currentColor" d="M26 26h-2.143v4.5H26c1.182 0 2.143-1.01 2.143-2.25S27.182 26 26 26z"/>`),
		g.Group(children),
	)
}