package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M220.6 121.2L271.1 96H448v96H333.2c-21.9-15.1-48.5-24-77.2-24s-55.2 8.9-77.2 24H64v-64h128c9.9 0 19.7-2.3 28.6-6.8zM0 128v288c0 35.3 28.7 64 64 64h384c35.3 0 64-28.7 64-64V96c0-35.3-28.7-64-64-64H271.1c-9.9 0-19.7 2.3-28.6 6.8L192 64h-32V48c0-8.8-7.2-16-16-16H80c-8.8 0-16 7.2-16 16v16C28.7 64 0 92.7 0 128zm168 176a88 88 0 1 1 176 0a88 88 0 1 1-176 0z"/>`),
		g.Group(children),
	)
}