package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ethernet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 224v192c0 17.7 14.3 32 32 32h64V336c0-8.8 7.2-16 16-16s16 7.2 16 16v112h64V336c0-8.8 7.2-16 16-16s16 7.2 16 16v112h64V336c0-8.8 7.2-16 16-16s16 7.2 16 16v112h64V336c0-8.8 7.2-16 16-16s16 7.2 16 16v112h64c17.7 0 32-14.3 32-32V224c0-17.7-14.3-32-32-32h-32v-32c0-17.7-14.3-32-32-32h-32V96c0-17.7-14.3-32-32-32H160c-17.7 0-32 14.3-32 32v32H96c-17.7 0-32 14.3-32 32v32H32c-17.7 0-32 14.3-32 32z"/>`),
		g.Group(children),
	)
}