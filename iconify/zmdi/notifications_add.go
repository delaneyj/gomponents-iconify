package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M150 408h84q0 18-12 30.5T192 451t-30-12.5t-12-30.5zm189-89l45 45v23H0v-23l45-45V195q0-52 32-91.5T158 52V37q0-14 10-24t24-10t24 10t10 24v15q49 12 81 51.5t32 91.5v124zm-62-81v-43h-64v-64h-42v64h-64v43h64v64h42v-64h64z"/>`),
		g.Group(children),
	)
}