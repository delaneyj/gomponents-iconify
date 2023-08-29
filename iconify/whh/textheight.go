package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Textheight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.74 257q-26 0-45-19t-19-45q0-25-42.5-44.5t-84.5-19.5h-1q-26 0-45 18.5t-19 45.5v639q0 27 19 45.5t45.5 18.5t45 19t18.5 45.5t-18.5 45t-45.5 18.5h-256q-26 0-45-18.5t-19-45t19-45.5t45.5-19t45-18.5t18.5-45.5V193q0-27-18.5-45.5t-45.5-18.5q-43 0-85.5 19.5t-42.5 44.5q0 26-18.5 45t-45 19t-45.5-19t-19-45V65q0-27 19-45.5t45-18.5h640q27 0 45.5 18.5t18.5 45.5v128q0 26-19 45t-45 19zm-776 640q8 8 8 19t-8 19l-66 82q-9 8-21.5 8t-21.5-8l-66-82q-9-8-9-19t9-19h56V129h-56q-9-8-9-19.5t9-19.5l66-82q9-8 21.5-8t21.5 8l66 82q8 8 8 19.5t-8 19.5h-55v768h55z"/>`),
		g.Group(children),
	)
}