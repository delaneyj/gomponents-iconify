package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FalloutShelter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 23C127.425 23 23 127.425 23 256s104.425 233 233 233s233-104.425 233-233S384.575 23 256 23zM149.268 71.287h213.168L256 256h213.17L362.732 440.713L256 256L149.564 440.713L42.83 256H256L149.268 71.287z"/>`),
		g.Group(children),
	)
}