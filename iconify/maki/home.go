package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2 13.748c0 .138.112.25.25.25h3.749v-3h3v3h3.749a.25.25 0 0 0 .25-.25v-5.75H2v5.75zm11.93-7.17l-.932-.82V2a1 1 0 1 0-2 0v2L7.681 1.09a.25.25 0 0 0-.353-.011l-.011.011l-6.25 5.463a.25.25 0 0 0 .18.42L3 7h10.747a.25.25 0 0 0 .183-.421z"/>`),
		g.Group(children),
	)
}