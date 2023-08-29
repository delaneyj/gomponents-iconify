package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VivoAppstore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.34"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".894" d="M32.343 11.785a8.637 8.637 0 0 1-16.686 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.807 33.637s8.218-4.807 17.181-.747c12.476 5.652 19.5-1.172 19.5-1.172"/>`),
		g.Group(children),
	)
}