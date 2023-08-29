package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.475.07h-1.486l-1.988 2.335v8.927c0 .751-1 .793-1 .043L12 8.148c.005-.094.027-.934-.582-1.574c-.368-.385-.822-.539-1.446-.585V.053L2 0v14.018h-.001a.012.012 0 0 0-.008.003H1v1.938h9.979v-1.938l-.991.054c-.006 0-.009-.003-.009-.003H9.97V7.006c.323.039.543.073.715.253c.327.336.316.842.314.867v3.25c0 1.082.729 1.604 1.562 1.604c.875 0 1.438-.604 1.438-1.647V2.959h.984v-.984l.373-1.023h1.037c.268 0 .563-.134.563-.399a.482.482 0 0 0-.481-.483zM2.958.939h6.073v5.092H2.958V.939zm6.125 6.977V9H6.958V7.916h2.125zM2.958 8H5.02v1H2.958V8z"/>`),
		g.Group(children),
	)
}