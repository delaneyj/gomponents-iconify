package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlizzardAuthenticator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.69 0 15.25-7.77 15.25-16.94v-20c-4 0-15.25-2-15.25-2s-11.26 2-15.25 2v20C8.75 35.73 22.31 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.39 18.9a3.39 3.39 0 1 0-6.78 0a3.33 3.33 0 0 0 2.07 3.1v3.13a1.32 1.32 0 0 0 2.64 0V22a3.39 3.39 0 0 0 2.07-3.1Z"/>`),
		g.Group(children),
	)
}