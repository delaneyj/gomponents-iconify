package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListOl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.983 19H4v-1h1.989v-.5h-.994v-1h.995V16H4v-1h2.983v4ZM21 18H9.069v-2H21v2ZM6.983 14H4v-.9L5.79 11H4v-1h2.983v.9L5.193 13h1.79v1ZM21 13H9.069v-2H21v2ZM6.486 9h-.995V6H4.5V5h1.986v4ZM21 8H9.069V6H21v2Z"/>`),
		g.Group(children),
	)
}