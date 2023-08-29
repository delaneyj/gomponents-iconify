package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Substitute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m0 5l.53-.53l2.25-2.25a.75.75 0 0 1 1.06 1.06l-.969.97h10.38a.75.75 0 0 1 0 1.5H2.87l.97.97a.75.75 0 0 1-1.06 1.06L.53 5.53L0 5Zm16 6l-.53-.53l-2.25-2.25a.75.75 0 1 0-1.061 1.06l.97.97H2.748a.75.75 0 0 0 0 1.5h10.38l-.97.97a.75.75 0 1 0 1.06 1.06l2.25-2.25L16 11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}