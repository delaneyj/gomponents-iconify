package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Googlecontaineroptimizedos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0zm0 21.6a9.6 9.6 0 0 1-5.016-1.416L11.28 17.7v-5.4L6.612 9.6v5.424l3.3 1.908l-4.152 2.4A9.6 9.6 0 0 1 7.296 3.6v4.8L12 11.136L16.68 8.4L12 5.724L8.688 7.632V2.964a9.6 9.6 0 0 1 12.372 5.64A9.72 9.72 0 0 1 21.672 12v.084L17.352 9.6l-4.68 2.712v5.412l4.68-2.7v-3.816l4.14 2.4A9.6 9.6 0 0 1 12 21.6z"/>`),
		g.Group(children),
	)
}