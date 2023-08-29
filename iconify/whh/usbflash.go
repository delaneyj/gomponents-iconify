package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usbflash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M987 860L860 987q-32 35-94 36.5t-118-20t-85-50.5L173 563q-13-13-13-32t13-33l325-325q14-14 33-14t32 14l390 390q29 29 50.5 85t20 118t-36.5 94zM0 307L307 0l154 154l-307 307zm292-138q6 6 15 6t15-6l31-31q7-6 7-15t-7-16l-31-30q-6-7-15-7t-15 7l-31 30q-6 7-6 16t6 15zM107 353q7 7 16 7t15-7l31-31q6-6 6-15t-6-15l-31-31q-6-6-15-6t-16 6l-30 31q-7 6-7 15t7 15z"/>`),
		g.Group(children),
	)
}