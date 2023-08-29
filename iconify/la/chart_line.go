package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m27.219 5.375l-3.5 4.375l-3.282-1.656l-.78-.375l-.5.718l-3.407 5.126l-3.156-2.376L12 10.75l-.594.438l-3.625 2.718l-3.531-.875l-.5 1.938l4 1l.469.125l.375-.282L12 13.25l3.406 2.563l.844.624l.594-.875l3.5-5.25l3.218 1.594l.72.344l4.5-5.625zm-7.157 12.938l-.843 1.062l-3.407 4.25l-3.218-2.438L12 20.75l-.594.438l-3.5 2.625l-3.468-1.72l-.875 1.813l4 2l.53.25l.5-.343L12 23.25l3.406 2.563l.781.562l.594-.75l3.125-3.906l3.25 4.843l.782 1.125l.843-1.062l4-5l-1.562-1.25l-3.125 3.906l-3.25-4.843z"/>`),
		g.Group(children),
	)
}