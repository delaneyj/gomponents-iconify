package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2.06a5.5 5.5 0 0 0-.5 10.97v8.41a.5.5 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-8.41A5.5 5.5 0 0 0 12 2.06Zm0 10a4.5 4.5 0 1 1 4.5-4.5a4.5 4.5 0 0 1-4.5 4.5Z"/>`),
		g.Group(children),
	)
}