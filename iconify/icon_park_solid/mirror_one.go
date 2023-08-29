package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMirrorOne0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="20" r="16" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 36v8m-10 0h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMirrorOne0)"/>`),
		g.Group(children),
	)
}