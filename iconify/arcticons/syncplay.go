package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Syncplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.073" cy="27.226" r="10.006" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.947 22.603S44.352 20.621 42.263 5.86a37.694 37.694 0 0 0-23.246 5.864C7.612 19.122 5.954 26.35 5.52 32.818c-.418 6.237 5.792 8.962 12.303 9.242c14.505 1.477 24.002-4.758 24.246-16.853c0 0-4.088 5.652-10.132 6.66"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.707 30.036c4.941-3.625 6.297-11.026 5.022-15.713M28.508 27.104l-4.38 2.53l-4.382 2.529V22.045l4.381 2.53Z"/>`),
		g.Group(children),
	)
}