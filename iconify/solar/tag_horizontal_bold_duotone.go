package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagHorizontalBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10.221 20h2.637c2.227 0 3.341 0 4.27-.501c.93-.502 1.52-1.42 2.701-3.259l.681-1.06C21.503 13.634 22 12.86 22 12c0-.86-.497-1.634-1.49-3.18l-.68-1.06c-1.181-1.838-1.771-2.757-2.701-3.259C16.199 4 15.085 4 12.858 4h-2.637C6.346 4 4.408 4 3.204 5.172C2 6.343 2 8.229 2 12c0 3.771 0 5.657 1.204 6.828C4.408 20 6.346 20 10.22 20Z" opacity=".5"/><path d="M7 7.055c.414 0 .75.316.75.706v8.475c0 .39-.336.706-.75.706s-.75-.316-.75-.706V7.76c0-.39.336-.706.75-.706Z"/></g>`),
		g.Group(children),
	)
}