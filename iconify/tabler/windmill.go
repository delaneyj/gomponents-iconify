package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windmill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12c2.76 0 5-2.01 5-4.5S14.76 3 12 3v9zm0 0c0 2.76 2.01 5 4.5 5s4.5-2.24 4.5-5h-9zm0 0c-2.76 0-5 2.01-5 4.5S9.24 21 12 21v-9zm0 0c0-2.76-2.01-5-4.5-5S3 9.24 3 12h9z"/>`),
		g.Group(children),
	)
}