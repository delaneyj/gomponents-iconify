package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m8.219 5l-.156.813l-.5 2.875l-.22 1.187H21.47l-.313 1.813H7.031l-.156.812l-.5 2.875l-.219 1.156h14.125l-.718 4.063L14.343 23l-4.093-2.25l.219-1.219l.219-1.156H5.811l-.125.813L5 23l-.156.75l.687.344l7.813 3.812l.406.188l.406-.157l9.156-3.843l.5-.188l.125-.562L27 6.187L27.219 5zm1.656 2h14.938l-2.75 15.469l-8.282 3.406l-6.687-3.25l.406-2.25h.781l-.25 1.438l.625.343L13.812 25l.438.25l.469-.219l6.156-2.843l.469-.22l.093-.53l1.032-5.72l.218-1.187H8.563l.126-.844h14.156l.125-.843l.687-3.813l.219-1.156H9.75z"/>`),
		g.Group(children),
	)
}