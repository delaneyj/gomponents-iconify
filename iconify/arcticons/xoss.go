package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xoss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.327 24.04L36.547 5.5H42.5v5.731L29.896 24.04L42.5 36.853V42.5h-5.7L18.326 24.04ZM11.117 5.5l10.245 10.236l-5.602 5.746L5.5 11.232V5.5h5.617ZM5.5 36.797V42.5h5.813l10.13-10.156l-5.784-5.788L5.5 36.797Z"/>`),
		g.Group(children),
	)
}