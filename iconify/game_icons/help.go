package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16C123.45 16 16 123.45 16 256s107.45 240 240 240s240-107.45 240-240S388.55 16 256 16zm0 60c99.41 0 180 80.59 180 180s-80.59 180-180 180S76 355.41 76 256S156.59 76 256 76zm0 30c-66.274 0-120 40.294-120 90c0 30 60 30 60 0c0-16.57 26.862-30 60-30c33.138 0 60 13.43 60 30s-30 15-60 30a19.594 19.594 0 0 0-4.688 3.28C226.53 244.986 226 271.926 226 286v15c0 16.62 13.38 30 30 30c16.62 0 30-13.38 30-30v-15c0-45 90-40.294 90-90s-53.726-90-120-90zm0 240a30 30 0 0 0-30 30a30 30 0 0 0 30 30a30 30 0 0 0 30-30a30 30 0 0 0-30-30z"/>`),
		g.Group(children),
	)
}