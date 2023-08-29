package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 10h2a1 1 0 0 1 0 2h-3a1 1 0 0 1-1-1V6a1 1 0 1 1 2 0v4zm7.63-1.562a9 9 0 1 1-17.26 0a5 5 0 1 1 7.668-6.387a9.102 9.102 0 0 1 1.924 0a5 5 0 1 1 7.668 6.387zm-.938-2.113a3 3 0 0 0-4.48-3.735a9.03 9.03 0 0 1 4.48 3.735zM6.787 2.59a3 3 0 0 0-4.48 3.735a9.03 9.03 0 0 1 4.48-3.735zM10 18a7 7 0 1 0 0-14a7 7 0 0 0 0 14z"/>`),
		g.Group(children),
	)
}