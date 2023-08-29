package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11c0 1.04-.156 2.045-.428 3H22v-8.023l-6 6.285l-6-6.285V19H5.428A10.95 10.95 0 0 1 5 16C5 9.935 9.935 5 16 5zm-4 10.969l4 4.187l4-4.187V21h5.785c-1.825 3.556-5.521 6-9.785 6s-7.96-2.444-9.785-6H12v-5.031z"/>`),
		g.Group(children),
	)
}