package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaHybridLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.508 40.763h32.984L27.008 17.407"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 35.802L27.008 7.237L12.883 31.702"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.992 7.237L4.5 35.802h27.732"/>`),
		g.Group(children),
	)
}