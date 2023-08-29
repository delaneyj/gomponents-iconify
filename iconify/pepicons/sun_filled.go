package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 10a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm-13 0a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2h-2a1 1 0 0 1-1-1Zm15 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1Zm-.086 4l1.414 1.414a1 1 0 0 1-1.414 1.414L14 15.414A1 1 0 0 1 15.414 14ZM3 3a1 1 0 0 1 1.414 0l1.414 1.414a1 1 0 1 1-1.414 1.414L3 4.414A1 1 0 0 1 3 3Zm7 12.5a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1Zm0-15a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1ZM16.828 3a1 1 0 0 1 0 1.414l-1.414 1.414A1 1 0 1 1 14 4.414L15.414 3a1 1 0 0 1 1.414 0Zm-11 11a1 1 0 0 1 0 1.414l-1.414 1.414A1 1 0 0 1 3 15.414L4.414 14a1 1 0 0 1 1.414 0Z"/>`),
		g.Group(children),
	)
}