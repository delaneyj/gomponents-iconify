package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uninstall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1.594l4.914 4.858l-1.406 1.422L13 5.394v7.11h-2v-7.11l-2.508 2.48l-1.406-1.422L12 1.594ZM2 2h5.5v2H4v10h16V4h-3.5V2H22v20H2V2Zm18 14H4v4h16v-4Zm-14.002.998h2.004v2.004H5.998v-2.004Zm3 0h2.004v2.004H8.998v-2.004Z"/>`),
		g.Group(children),
	)
}