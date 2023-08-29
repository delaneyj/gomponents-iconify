package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emailtrace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1017.084 734l-59-59q2-19 2-35q0-110-67.5-196t-171.5-113l296-296q7 14 7 29v640q0 15-7 30zm-377-414q-106 0-192 65l-385-385h898l-320 320h-1zm-633 414q-7-15-7-30V64q0-15 7-29l350 349zm313-94q0 67 27 128h-284l325-324q-68 87-68 196zm320-256q106 0 181 75t75 181q0 75-41 138l161 160q8 8 8 19.5t-8 19.5l-39 39q-8 8-19.5 8t-19.5-8l-160-160q-64 40-138 40q-106 0-181-75t-75-181t75-181t181-75zm0 384q53 0 90.5-37.5t37.5-90.5t-37.5-90.5t-90.5-37.5t-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5z"/>`),
		g.Group(children),
	)
}