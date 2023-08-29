package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.898 6.135a3.501 3.501 0 0 0-4.884.808l-8.832 12.33a2.5 2.5 0 1 0 4.064 2.91l6.236-8.705l-.813-.583l-6.235 8.707a1.502 1.502 0 0 1-2.092.347a1.502 1.502 0 0 1-.345-2.094l8.83-12.33h.002l-.002-.002a2.506 2.506 0 0 1 3.49-.576a2.506 2.506 0 0 1 .576 3.49v-.002l-9.68 13.516a3.506 3.506 0 0 1-4.884.81a3.507 3.507 0 0 1-.807-4.886l7.035-9.822l-.813-.582l-7.035 9.822a4.497 4.497 0 0 0 1.04 6.277a4.496 4.496 0 0 0 6.277-1.036l9.68-13.516a3.501 3.501 0 0 0-.81-4.883z"/>`),
		g.Group(children),
	)
}