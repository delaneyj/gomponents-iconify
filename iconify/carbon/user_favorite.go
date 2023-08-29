package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserFavorite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.303 12a2.662 2.662 0 0 0-1.908.806l-.393.405l-.397-.405a2.662 2.662 0 0 0-3.816 0a2.8 2.8 0 0 0 0 3.896L25.002 21l4.209-4.298a2.8 2.8 0 0 0 0-3.896A2.662 2.662 0 0 0 27.303 12zM2 30h2v-5a5.006 5.006 0 0 1 5-5h6a5.006 5.006 0 0 1 5 5v5h2v-5a7.008 7.008 0 0 0-7-7H9a7.008 7.008 0 0 0-7 7zM12 4a5 5 0 1 1-5 5a5 5 0 0 1 5-5m0-2a7 7 0 1 0 7 7a7 7 0 0 0-7-7z"/>`),
		g.Group(children),
	)
}