package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m22 3.586l-4.521 4.521l-6.74 1.928a4.061 4.061 0 0 0-2.802 2.65L3.86 25.274l1.434 1.434l1.434 1.434l12.593-4.08a4.057 4.057 0 0 0 2.64-2.788l1.929-6.748L28.414 10L22 3.586zm0 2.828L25.586 10L23 12.586L19.414 9L22 6.414zm-4.29 3.711l4.165 4.164l-1.842 6.45a2.064 2.064 0 0 1-1.336 1.421L7.69 25.725l5.795-5.795A2 2 0 0 0 14 20a2 2 0 0 0 0-4a2 2 0 0 0-1.93 2.516L6.275 24.31l3.563-11a2.072 2.072 0 0 1 1.437-1.348l6.436-1.838z"/>`),
		g.Group(children),
	)
}