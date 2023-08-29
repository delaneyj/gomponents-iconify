package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M768 96v1088q0 26-19 45t-45 19t-45-19L326 896H64q-26 0-45-19T0 832V448q0-26 19-45t45-19h262L659 51q19-19 45-19t45 19t19 45z"/>`),
		g.Group(children),
	)
}