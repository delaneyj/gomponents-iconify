package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapWizardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.16 42.5l-18.5-17.179V5.5H8.805"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.66 25.321H10.126v-2.643"/><circle cx="13.429" cy="15.411" r="5.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="5.5" cy="9.464" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.16 21.357v1.321"/>`),
		g.Group(children),
	)
}