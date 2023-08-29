package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Solarpanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M990.86 703h-281l58 288q3 16-7 24t-25 8h-448q-15 0-24.5-8t-7.5-24l58-288h-281q-16 0-25.5-7.5T.86 671l127-639q2-12 8.5-22t15.5-10h709q10 0 21.5 11t13.5 21l127 639q3 17-6.5 24.5t-25.5 7.5zm-159-639h-640l-96 575h831zm-484 320l-28 191h-160l41-191h147zm264 0l28 191h-256l27-191h201zm210 0l42 191h-160l-28-191h146zm-155-64l-27-192h128l41 192h-142zm-246 0l27-192h128l27 192h-182zm-206 0l41-192h128l-27 192h-142z"/>`),
		g.Group(children),
	)
}