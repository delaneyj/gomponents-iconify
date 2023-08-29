package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCameraOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="19" r="14"/><circle cx="24" cy="19" r="6" fill="#fff"/><path d="m17 31l-6 12h26l-6-12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCameraOne0)"/>`),
		g.Group(children),
	)
}