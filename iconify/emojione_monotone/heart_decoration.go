package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartDecoration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32c0 16.568 13.432 30 30 30s30-13.432 30-30C62 15.432 48.568 2 32 2m0 48c-1.371-1.814-20.53-12.883-16.602-25.218c3.53-11.073 15.094-6.597 16.602-.594c1.094-5.635 12.949-10.694 16.604.584C52.529 36.908 33.367 48.557 32 50"/>`),
		g.Group(children),
	)
}