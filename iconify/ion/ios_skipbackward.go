package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosSkipbackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M96 96v320h79V274.2L416 416V96L175 237.8V96H96z" fill="currentColor"/>`),
		g.Group(children),
	)
}