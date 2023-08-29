package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnifiedRemote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="arcticonsUnifiedRemote0" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 5.5A31.5 31.5 0 0 1 42.5 37"/></defs><use href="#arcticonsUnifiedRemote0" stroke-linecap="round" stroke-linejoin="round"/><use href="#arcticonsUnifiedRemote0" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11" cy="37" r="7.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.392 5.835A31.527 31.527 0 0 1 42.5 37a31.766 31.766 0 0 1-.335 4.608"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.713 42.5A19.516 19.516 0 0 0 5.5 18.287"/>`),
		g.Group(children),
	)
}