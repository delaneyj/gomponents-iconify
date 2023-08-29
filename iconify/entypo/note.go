package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Note(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.971 9.438c-.422.656-.646.375-.52 0c.336-.993.348-4.528-2.451-4.969L11.998 16c0 1.657-1.735 4-4.998 4c-1.657 0-3-.871-3-2.5c0-2.119 1.927-3.4 4-3.4c1.328 0 2 .4 2 .4V0h2c0 2.676 5.986 4.744 2.971 9.438z"/>`),
		g.Group(children),
	)
}