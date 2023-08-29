package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m896 320l64 576q0 53-37.5 90.5T832 1024t-90.5-37.5T704 896l64-576q-31 0-61-36t-48.5-88T640 96q0-21 1-33t8.5-30T672 0h32v64q0 54 37 91t91 37t91-37t37-91V0h32q15 15 22.5 33t8.5 30t1 33q0 48-18.5 100T957 284t-61 36zm-63.5 512q-26.5 0-45.5 18.5T768 896t19 45.5t45.5 18.5t45-19t18.5-45.5t-18.5-45t-45-18.5zM256 608q0 40 18.5 68t45.5 28v160q0 66-47 113t-113 47t-113-47T0 864V704q27 0 45.5-28T64 608t-18.5-68T0 512q0-27 19-45.5T64 448h64V256l-64-64L96 0h128l32 192l-64 64v192h64q27 0 45.5 18.5T320 512q-26 0-45 28t-19 68z"/>`),
		g.Group(children),
	)
}