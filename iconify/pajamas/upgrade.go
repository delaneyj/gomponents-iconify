package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upgrade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m14 6l-1.5-1.5l-3.44-3.44L8 0L6.94 1.06L3.5 4.5L2 6h3v7h6V6h3ZM8 2.121L10.379 4.5H9.5v7h-3v-7h-.879L8 2.121ZM5.75 14.5a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5h-4.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}