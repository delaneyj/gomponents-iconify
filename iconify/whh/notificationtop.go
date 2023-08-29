package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notificationtop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M855 1024H169q-57 0-113-57T0 852V428q0-58 56-115t113-57h151L512 0l192 256h151q57 0 113 57t56 115v424q0 58-56 115t-113 57zm-23-576H192q-26 0-45 19t-19 45t19 45t45 19h640q26 0 45-19t19-45t-19-45t-45-19zm0 256H192q-26 0-45 18.5t-19 45t19 45.5t45 19h640q26 0 45-19t19-45.5t-19-45t-45-18.5z"/>`),
		g.Group(children),
	)
}