package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSlovakia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 21H16v11.91c0 5.606 4.257 8.569 8 10.09c3.743-1.521 8-4.484 8-10.09V21m-3 10h-4v6h-2v-6h-4v-2h4v-2h-2v-2h2v-2h2v2h2v2h-2v2h4v2"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm-3.085 41c-1.546 1.182-3.269 2.08-4.916 2.75c-1.646-.67-3.369-1.568-4.915-2.75H6.254a27.943 27.943 0 0 1-.679-1.756h11.572C15.287 39.219 14 36.578 14 33.137V22H5.854C9.888 11.486 20.083 4 32 4s22.112 7.486 26.147 18H34v11.137c0 3.441-1.287 6.082-3.147 8.107h27.572A27.943 27.943 0 0 1 57.746 43H28.915z"/>`),
		g.Group(children),
	)
}