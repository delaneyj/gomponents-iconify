package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Studysmarter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.578 8.944l17.395 10.055a1.054 1.054 0 0 1 0 1.825L25.578 30.879A3.16 3.16 0 0 0 24 33.614v4.532a1.052 1.052 0 0 0 1.578.91l7.27-4.202a3.16 3.16 0 0 0 1.579-2.735v-4.633a3.16 3.16 0 0 0-1.579-2.736l-7.27-4.202A3.16 3.16 0 0 1 24 17.811V9.854a1.052 1.052 0 0 1 1.578-.91Zm-3.156 0L5.027 18.999a1.054 1.054 0 0 0 0 1.825l17.395 10.055A3.16 3.16 0 0 1 24 33.614v4.532a1.052 1.052 0 0 1-1.578.91l-7.27-4.202a3.16 3.16 0 0 1-1.579-2.735v-4.633a3.16 3.16 0 0 1 1.579-2.736l7.27-4.202A3.16 3.16 0 0 0 24 17.811V9.854a1.052 1.052 0 0 0-1.578-.91Z"/>`),
		g.Group(children),
	)
}