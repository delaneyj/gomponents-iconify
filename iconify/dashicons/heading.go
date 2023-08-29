package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.5 4v5.2h-5V4H5v13h2.5v-5.2h5V17H15V4z"/>`),
		g.Group(children),
	)
}