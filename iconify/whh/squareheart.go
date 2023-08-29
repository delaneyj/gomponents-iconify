package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareheart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-224-832q-66 0-113 47t-47 113q0-66-47-113t-113-47t-113 47t-47 113q0 80 31 148.5t105 143.5l184 188l183-188q75-76 106-144t31-148q0-66-47-113t-113-47z"/>`),
		g.Group(children),
	)
}