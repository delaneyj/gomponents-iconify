package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.858 20h-2.637c-3.875 0-5.813 0-7.017-1.172C2 17.657 2 15.771 2 12c0-3.771 0-5.657 1.204-6.828C4.408 4 6.346 4 10.22 4h2.637c2.227 0 3.341 0 4.27.501c.93.502 1.52 1.42 2.701 3.259l.681 1.06C21.503 10.366 22 11.14 22 12c0 .86-.497 1.634-1.49 3.18l-.68 1.06c-1.181 1.838-1.771 2.757-2.701 3.259c-.93.501-2.044.501-4.271.501ZM7 7.054c.414 0 .75.316.75.706v8.475c0 .39-.336.706-.75.706s-.75-.316-.75-.706V7.76c0-.39.336-.706.75-.706Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}