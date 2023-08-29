package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m981.356 426l-548 79q-15 10-30 4l-19 3v-14l-256-153v615q0 26-18.5 45t-45 19t-45.5-19t-19-45V64q0-27 19-45.5T64.856 0t45 18.5t18.5 45.5v103l256-153V0l19 3q15-6 30 4l548 79q18 2 30.5 15.5t12.5 31.5v246q0 18-12.5 31.5t-30.5 15.5zm-853-184v28l256 152V90zm384-137l-64-9v320l64-9V105zm256 34l-128-17v268l128-17V139zm192 57q0-14-9.5-24t-23.5-12l-31-4v200l31-4q14-2 23.5-12t9.5-24V196z"/>`),
		g.Group(children),
	)
}