package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hanger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 640v8q167 22 275.5 91T1024 896q0 26-28 45t-68 19t-68-19t-28-45q0-35-43-64.5T672.5 785T512 768t-160.5 17T235 831.5T192 896q0 26-28 45t-68 19t-68-19t-28-45q0-89 108.5-157.5T384 648v-8q0-26 19-45t45-19V448q0-26 19-45t45-19q53 0 90.5-37.5T640 256t-37.5-90.5T512 128t-90.5 37.5T384 256q0 26-18.5 45t-45 19t-45.5-19t-19-45q0-106 75-181T512 0t181 75t75 181q0 88-54 157t-138 91v72q27 0 45.5 19t18.5 45z"/>`),
		g.Group(children),
	)
}