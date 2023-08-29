package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm12-13.038L20.516 12.5h-2.375l4.625 3.998l-1.461 1.257l-3.039-3.708h-2.01l-3.652 4.445l-.398.486l1.788-4.923H12.37l-1.184 3.325L10 14.056H8l-1.188 3.332l-1.192-3.346H4l1.773 4.911h2.112L9 15.823l1.124 3.155h2.08L10.96 20.5h1.997l4.305-5.241l1.233 1.508h-1.619l-.833 1.029h3.292l.955 1.16h1.998l1.671-1.43l1.677 1.436zm-2.345-2.972l2.249-1.902l-2.251-.002l-1.113.948z"/>`),
		g.Group(children),
	)
}