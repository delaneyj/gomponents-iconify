package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCameraThree0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><rect width="36" height="26" x="6" y="14" stroke-linejoin="round" rx="2"/><path fill="#fff" stroke-linejoin="round" d="m10 14l2.167-6h7.666L22 14H10Z"/><circle cx="29" cy="27" r="7" fill="#fff" stroke-linejoin="round"/><path d="M36 10v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCameraThree0)"/>`),
		g.Group(children),
	)
}