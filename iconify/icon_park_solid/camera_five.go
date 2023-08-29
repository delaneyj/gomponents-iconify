package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCameraFive0"><g fill="none" stroke-width="4"><circle cx="24" cy="21" r="16" fill="#fff" stroke="#fff"/><circle cx="24" cy="21" r="7" fill="#000" stroke="#000"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M16 43h16m-8-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCameraFive0)"/>`),
		g.Group(children),
	)
}