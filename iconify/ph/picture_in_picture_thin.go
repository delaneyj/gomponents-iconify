package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 52H40a12 12 0 0 0-12 12v128a12 12 0 0 0 12 12h176a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12ZM36 192V64a4 4 0 0 1 4-4h176a4 4 0 0 1 4 4v60h-76a12 12 0 0 0-12 12v60H40a4 4 0 0 1-4-4Zm180 4h-76v-60a4 4 0 0 1 4-4h76v60a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}