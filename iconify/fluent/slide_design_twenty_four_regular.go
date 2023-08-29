package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideDesignTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M16.25 3.75a.75.75 0 0 1 .75.75V10h2.5a.75.75 0 0 1 0 1.5H17v5.55a2.512 2.512 0 0 0-1.5.158V4.5a.75.75 0 0 1 .75-.75zM16.5 18a1.499 1.499 0 1 1 0 3a1.5 1.5 0 0 1 0-3zM4 5.5a.75.75 0 0 1 .75-.75h7a.75.75 0 0 1 .75.75c0 1.744-.362 4.369-1.48 6.588c-1.13 2.243-3.096 4.162-6.27 4.162a.75.75 0 0 1 0-1.5c2.426 0 3.96-1.414 4.93-3.338c.837-1.66 1.206-3.632 1.297-5.162H4.75A.75.75 0 0 1 4 5.5zM12.5 21a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3zm9.5-1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}