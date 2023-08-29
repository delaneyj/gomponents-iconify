package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 103.564v792.188h435.718v114.186H329.517v86.498h540.967v-86.498H764.282V895.752H1200V103.564H0zm147.583 149.561h904.833v493.14H147.583v-493.14z"/>`),
		g.Group(children),
	)
}