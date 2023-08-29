package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3c-1.864 0-3.4 1.275-3.844 3H3v20h26V6h-9.156C19.4 4.275 17.864 3 16 3zm0 2c.81 0 1.428.385 1.75 1h-3.5c.322-.615.94-1 1.75-1zM5 8h22v9H5V8zm11 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM5 19h22v5H5v-5z"/>`),
		g.Group(children),
	)
}