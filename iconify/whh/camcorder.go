package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camcorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928.428 640h-320q-40 0-68-28t-28-68v-32h-64v384q0 53-37.5 90.5t-90.5 37.5h-192q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h192q53 0 90.5 37.5t37.5 90.5v192h64v-32q0-40 28-68t68-28h320q40 0 68 28t28 68v256q0 40-28 68t-68 28zm-832-64q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5zm128-448q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113t-113-47zm0 256q-40 0-68-28t-28-68t28-68t68-28t68 28t28 68t-28 68t-68 28z"/>`),
		g.Group(children),
	)
}