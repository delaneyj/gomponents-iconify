package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cleaner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.732 38.242c4.055 3.835 7.865 4.258 7.865 4.258c3.178-2.839 4.91-5.375 6.842-10.807c0 0-7.548-7.321-9.091-9.625c0 0-3.509 1.8-6.223 2.048c-2.046.187-5.536-.86-5.536-.86c.028 2.918 2.023 11.088 6.143 14.986Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.348 22.068c1.903-1.076 5.383-3.554 6.994-2.15a26.204 26.204 0 0 1 3.477 3.421c1.099 1.434-.418 5.077-1.38 8.354"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.055 21.448C29.422 18.428 42.411 5.5 42.411 5.5M7.03 30.106c4.237.317 7.03-.612 7.03-.612m-3.142 7.893c6.19-.867 8.333-3.97 8.333-3.97"/>`),
		g.Group(children),
	)
}