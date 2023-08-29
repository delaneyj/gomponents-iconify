package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creativecommons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-896q-104 0-192.5 51.5t-140 140T128 512t51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512t-51.5-192.5t-140-140T512 128zm128 320q-27 0-45.5 18.5T576 512t18.5 45.5T640 576h110q-17 29-46 46.5T640 640q-53 0-90.5-37.5T512 512t37.5-90.5T640 384q35 0 64 17.5t46 46.5H640zm-320 64q0 26 19 45t45 19h110q-17 29-46 46.5T384 640q-53 0-90.5-37.5T256 512t37.5-90.5T384 384q35 0 64 17.5t46 46.5H384q-26 0-45 19t-19 45z"/>`),
		g.Group(children),
	)
}