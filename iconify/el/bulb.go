package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 119.608c-158.775 0-287.486 96.415-287.486 215.368l163.322 494.645h248.328l163.322-494.659c0-118.951-128.702-215.369-287.486-215.369v.015zm0 46.996c158.775 0 240.479 128.138 240.479 168.362L724.161 680.393H475.833L359.515 334.966c0-49.63 81.704-168.362 240.479-168.362H600zM472.215 865.699v85.982h255.57v-85.984h-255.57v.002zm0 128.725v85.982h255.57v-85.982h-255.57z"/>`),
		g.Group(children),
	)
}