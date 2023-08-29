package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orkut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M232 2Q137 2 69.5 69.5T2 232t67.5 162.5T232 462t162.5-67.5T462 232T394.5 69.5T232 2zm0 387q-65 0-111-46T75 232t46-111t111-46t111 46t46 111t-46 111t-111 46z"/>`),
		g.Group(children),
	)
}