package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForDjibouti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M15.935 30.279L14.3 25.117l-1.588 5.162H7.417l4.235 3.352l-1.604 5.252l4.252-3.229l4.251 3.229l-1.603-5.252l4.236-3.352z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 31L12.934 52.48C7.443 47.365 4 40.078 4 32c0-8.079 3.443-15.365 8.934-20.48L32 31h28a55.56 55.56 0 0 1 0 2H32z"/>`),
		g.Group(children),
	)
}