package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M480.48 480.465L240.24 0L0 480.465h208.705V512h63.07v-31.535H480.48zm-208.705-63.07V204.088l106.647 213.305H271.775zm-169.717 0l106.647-213.306v213.305H102.058z"/>`),
		g.Group(children),
	)
}