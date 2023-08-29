package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DdfGuide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="9.539" cy="32.9" r=".75" fill="currentColor"/><circle cx="24" cy="32.9" r=".75" fill="currentColor"/><circle cx="38.462" cy="32.9" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 19.388c0-1.374.458-2.748 1.374-3.664c.916-.916 2.061-1.374 3.665-1.374c2.748 0 5.038 2.29 5.038 5.038c0 1.374-.458 2.748-1.374 3.664c-1.145.917-3.664 2.29-3.664 4.352m9.423-8.016c0-1.374.458-2.748 1.374-3.664c.916-.916 2.06-1.374 3.664-1.374c2.748 0 5.038 2.29 5.038 5.038c0 1.374-.458 2.748-1.374 3.664C26.52 23.97 24 25.343 24 27.405m9.423-8.017c0-1.374.458-2.748 1.374-3.664c.916-.916 2.061-1.374 3.665-1.374c2.748 0 5.038 2.29 5.038 5.038c0 1.374-.458 2.748-1.374 3.664c-1.145.917-3.664 2.29-3.664 4.352"/>`),
		g.Group(children),
	)
}