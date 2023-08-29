package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.78 11.16a.75.75 0 0 0-1.06 0l-1.97 1.97V2.87l1.97 1.97a.75.75 0 1 0 1.06-1.06L8.53.53L8 0l-.53.53l-3.25 3.25a.75.75 0 0 0 1.06 1.06l1.97-1.97v10.26l-1.97-1.97a.75.75 0 0 0-1.06 1.06l3.25 3.25L8 16l.53-.53l3.25-3.25a.75.75 0 0 0 0-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}