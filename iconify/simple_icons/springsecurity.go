package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Springsecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.59 2.066L11.993 0L3.41 2.066v6.612h4.557a3.804 3.804 0 0 0 0 .954H3.41v3.106C3.41 19.867 11.994 24 11.994 24s8.582-4.133 8.582-11.258V9.635h-4.545a3.616 3.616 0 0 0 0-.954h4.558zM12 12.262h-.006a3.109 3.109 0 1 1 .006 0zm-.006-4.579a.804.804 0 0 0-.37 1.52v.208l.238.237v.159l.159.159v.159l-.14.14l.15.246v.159l-.16.189l.223.222l.246-.246V9.218a.804.804 0 0 0-.346-1.535zm0 .836a.299.299 0 1 1 .298-.299a.299.299 0 0 1-.298.3z"/>`),
		g.Group(children),
	)
}