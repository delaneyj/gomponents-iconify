package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moovit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5s15.496-12.013 15.496-23.504a15.496 15.496 0 0 0-30.992 0C8.504 31.487 24 43.5 24 43.5"/><circle cx="17.203" cy="18.555" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.75 22.322A8.01 8.01 0 0 1 24 24a8.01 8.01 0 0 1-4.75-1.678"/><circle cx="30.797" cy="18.555" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}