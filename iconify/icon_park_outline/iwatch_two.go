package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IwatchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="4" d="M15.417 10.5C18.237 7.7 21.942 6 26 6c8.837 0 16 8.059 16 18s-7.163 18-16 18c-4.058 0-7.763-1.7-10.583-4.5"/><rect width="10" height="28" x="6" y="10" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><rect width="4" height="4" x="13" y="18" fill="currentColor" rx="2" transform="rotate(90 13 18)"/><rect width="4" height="4" x="13" y="25" fill="currentColor" rx="2" transform="rotate(90 13 25)"/></g>`),
		g.Group(children),
	)
}