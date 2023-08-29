package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M768 96v1088q0 26-19 45t-45 19t-45-19L326 896H64q-26 0-45-19T0 832V448q0-26 19-45t45-19h262L659 51q19-19 45-19t45 19t19 45zm384 544q0 76-42.5 141.5T997 875q-10 5-25 5q-26 0-45-18.5T908 816q0-21 12-35.5t29-25t34-23t29-36t12-56.5t-12-56.5t-29-36t-34-23t-29-25t-12-35.5q0-27 19-45.5t45-18.5q15 0 25 5q70 27 112.5 93t42.5 142z"/>`),
		g.Group(children),
	)
}