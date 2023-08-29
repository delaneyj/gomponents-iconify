package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VuejsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.018 19.151L4.315 6h2.823l4.886 8.407L16.871 6h2.809z" opacity=".5"/><path fill="currentColor" d="m14.38 4.001l-2.374 3.956l-2.384-3.956H.825L12.02 23.115L23.161 4H14.38zm-2.362 15.15L4.315 6h2.823l4.886 8.407L16.871 6h2.809l-7.662 13.151z"/>`),
		g.Group(children),
	)
}