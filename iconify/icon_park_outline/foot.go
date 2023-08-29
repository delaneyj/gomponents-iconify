package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 8c1.766 7.879 2.783 24.837-5.693 28.618c-.481.316-2.15.851-5.618.473c-2.006-.219-4.284 2.263-1.123 5.673c.776.838 2.504 1.702 6.741.945h7.705c1.926 0 5.105-2.363 8.186-8.036c1.124-1.576 4.334-4.822 8.187-5.2c2.408-.158 6.838-2.08 5.297-8.51C39.926 15.506 33.54 9.296 32 4"/>`),
		g.Group(children),
	)
}