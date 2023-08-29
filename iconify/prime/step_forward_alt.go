package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepForwardAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 4.25c-.41 0-.75.34-.75.75v6.19L8.53 4.47A.751.751 0 0 0 7.25 5v14c0 .3.18.58.46.69a.75.75 0 0 0 .82-.16l6.72-6.72V19c0 .41.34.75.75.75s.75-.34.75-.75V5c0-.41-.34-.75-.75-.75ZM8.75 17.19V6.81L13.94 12l-5.19 5.19Z"/>`),
		g.Group(children),
	)
}