package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M591 526h126v65H591v126h-66V591H126V192H0v-66h126V0h65v126h374l74-73l35 35l-83 83v355zm-400-26l309-308H191v308zm334-263L237 526h288V237z"/>`),
		g.Group(children),
	)
}