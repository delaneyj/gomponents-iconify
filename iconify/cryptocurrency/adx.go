package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM11.587 8.602L9 11.36l4.36 4.614L9 20.68l2.587 2.717L16 18.792l4.305 4.604L23 20.603l-4.396-4.604L23 11.36l-2.695-2.757L16 13.208l-4.413-4.606zm.355-.376l2.126 2.265l1.914-2.114l1.914 2.114l2.126-2.265L15.982 4l-4.04 4.226zm0 15.548L15.982 28l4.04-4.226l-2.126-2.265l-1.914 2.114l-1.914-2.114l-2.126 2.265z"/>`),
		g.Group(children),
	)
}