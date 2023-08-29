package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notificationbottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M855.048 768h-151l-192 256l-192-256h-151q-57 0-113-57t-56-115V172q0-58 56-115t113-57h686q57 0 113 57t56 115v424q0 58-56 115t-113 57zm-23-576h-640q-26 0-45 19t-19 45.5t19 45t45 18.5h640q26 0 45-18.5t19-45t-19-45.5t-45-19zm0 256h-640q-26 0-45 19t-19 45.5t19 45t45 18.5h640q26 0 45-18.5t19-45t-19-45.5t-45-19z"/>`),
		g.Group(children),
	)
}