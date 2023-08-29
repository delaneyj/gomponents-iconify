package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Price(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1006.914 680l-325 325q-18 18-44.5 18t-44.5-18l-574-574q-18-18-18-79V64q0-27 19-45.5t45-18.5h288q62 0 80 18l574 573q18 19 18 45t-18 44zm-750-553q-53 0-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5t90.5-37.5t37.5-90.5t-37.5-90.5t-90.5-37.5z"/>`),
		g.Group(children),
	)
}