package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m517 480l74-106v303H0V87h476l-51 74H74v442l443 1V480zm66-449L347 372L182 250L99 365l280 214l321-461z"/>`),
		g.Group(children),
	)
}