package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.28 5.84a.75.75 0 0 1-1.06-1.06l3.25-3.25L8 1l.53.53l3.25 3.25a.75.75 0 0 1-1.06 1.06L8.75 3.87V12a1.5 1.5 0 0 0 1.5 1.5h4a.75.75 0 0 1 0 1.5h-4a3 3 0 0 1-3-3V3.87L5.28 5.84Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}