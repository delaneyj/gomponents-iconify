package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminalalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-406-562l-187-188q-19-20-46.5-20t-47 19.5t-19.5 47t20 46.5l141 143l-141 142q-20 20-20 47.5t19.5 46.5t47 19t46.5-19l187-188q18-17 19-42t-14-44q2-3-5-10zm278 177h-192q-26 0-45 19t-19 45.5t19 45.5t45 19h192q27 0 45.5-19t18.5-45.5t-18.5-45.5t-45.5-19z"/>`),
		g.Group(children),
	)
}