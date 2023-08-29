package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm9.992-17.529L26 14.47l-3.826-6.462l.002-.007H9.922l.002.008l-.002-.002L6 14.563l.039.006l-.032.013L16.097 26l.426-.52zm-13.136.459l6.423-.011l2.122 3.635l-5.363 6.162l-5.353-6.112zm-2.778 2.98l-2.602-2.972l4.361-.007zm9.723-3.846l-3.038-5.205l4.883-.01l3.11 5.207zm4.771.846l-2.56 2.942l-1.713-2.935zm-5.791-.844l-5.42.009l2.729-4.62zm-8.399-5.194l5.04-.01l-3.08 5.214l-5.06.009z"/>`),
		g.Group(children),
	)
}