package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m8.175 13l-3.324 6h11.824L20 13H8.175zM5.865 2.299L0 12.914l3.372 6.084L9.238 8.383L5.865 2.299zM19.445 12L13.349 1H6.604l6.088 11h6.753z"/>`),
		g.Group(children),
	)
}