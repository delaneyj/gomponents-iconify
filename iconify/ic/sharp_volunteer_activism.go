package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpVolunteerActivism(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 11h4v11H1zm15-7.75C16.65 2.49 17.66 2 18.7 2C20.55 2 22 3.45 22 5.3c0 2.27-2.91 4.9-6 7.7c-3.09-2.81-6-5.44-6-7.7C10 3.45 11.45 2 13.3 2c1.04 0 2.05.49 2.7 1.25zM22 17h-9l-2.09-.73l.33-.95L13 16h4v-2l-8.03-3H7v9.02L14 22l8-3z"/>`),
		g.Group(children),
	)
}