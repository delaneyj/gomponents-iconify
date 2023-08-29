package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 832q-40 0-68-28t-28-68V416q0-42 45-101t83-59q19 0 32.5 17t19.5 44t9 51.5t3 47.5v320q0 40-28 68t-68 28zM704 384H320q-27 0-45.5-19T256 320L192 64q0-27 18.5-45.5T256 0h512q26 0 45 19t19 45l-64 256q0 26-19 45t-45 19zM192 736q0 40-28 68t-68 28t-68-28t-28-68V416q0-23 2.5-47.5t9-51.5t20-44T64 256q38 0 83 59t45 101v320zm128-288h384q26 0 45 18.5t19 45t-19 45.5t-45 19H320q-27 0-45.5-19T256 511.5t18.5-45T320 448zm0 192h384q26 0 45 18.5t19 45.5v64q0 26-19 45t-45 19H320q-27 0-45.5-19T256 768v-64q0-27 18.5-45.5T320 640z"/>`),
		g.Group(children),
	)
}