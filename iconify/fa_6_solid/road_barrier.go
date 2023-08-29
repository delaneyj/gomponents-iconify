package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadBarrier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 32C14.3 32 0 46.3 0 64v384c0 17.7 14.3 32 32 32s32-14.3 32-32V266.3L149.2 96H64V64c0-17.7-14.3-32-32-32zm373.2 64h-74.4l-5.4 10.7L234.8 288h74.3l5.4-10.7L405.2 96zm-42.4 192h74.3l5.4-10.7L533.2 96h-74.4l-5.4 10.7L362.8 288zm-160-192l-5.4 10.7L106.8 288h74.3l5.4-10.7L277.2 96h-74.4zm288 192H576v160c0 17.7 14.3 32 32 32s32-14.3 32-32V64c0-17.7-14.3-32-32-32s-32 14.3-32 32v53.7L490.8 288z"/>`),
		g.Group(children),
	)
}