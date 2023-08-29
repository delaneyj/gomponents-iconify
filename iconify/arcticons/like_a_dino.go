package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LikeADino(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.196 18.172c.459-10.217 19.877-16.964 27.151 2.157c6.566 18.864-16.426 22.457-29.754 17.41C2.71 35.079.033 21.554 15.196 18.17h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.752 29.939s6.322 6.235 13.053 1.163"/><circle cx="20.639" cy="16.388" r=".75" fill="currentColor"/><circle cx="28.71" cy="17.604" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}