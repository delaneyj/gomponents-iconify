package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func External(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M501 107L167 443l70 71l335-335l62 61c17 17 32 12 32-14V45c0-16-14-31-31-31H454c-26 0-32 15-15 32zm37 465V345h76v253c0 45-37 82-82 82H82c-45 0-82-37-82-82V147c0-45 37-82 82-81h277v76H108c-16 0-31 15-31 31v399c0 16 15 31 31 31h399c16 0 31-15 31-31z"/>`),
		g.Group(children),
	)
}