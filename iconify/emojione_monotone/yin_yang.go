package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="42.264" cy="42.263" r="5.443" fill="currentColor"/><path fill="currentColor" d="M32 2C15.458 2 2 15.458 2 32s13.458 30 30 30s30-13.458 30-30S48.542 2 32 2m-6.416 23.584a5.444 5.444 0 1 1-7.699-7.7a5.444 5.444 0 0 1 7.699 7.7m20.501 30.675c-4.859 1.321-10.27.086-14.086-3.729c-5.668-5.668-5.668-14.86 0-20.529c5.669-5.669 5.669-14.86 0-20.528c-3.815-3.816-9.225-5.052-14.084-3.73A27.875 27.875 0 0 1 32 3.936c15.476 0 28.064 12.589 28.064 28.064c0 10.344-5.628 19.391-13.979 24.259"/>`),
		g.Group(children),
	)
}