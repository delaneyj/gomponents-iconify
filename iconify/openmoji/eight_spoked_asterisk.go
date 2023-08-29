package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightSpokedAsterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 20.981v30.038M51.019 36H20.981M46.62 25.38L25.38 46.62m21.24 0L25.38 25.38"/><path fill="#5c9e31" d="M60 60.958H12a.945.945 0 0 1-1-1v-48a.945.945 0 0 1 1-1h48a.945.945 0 0 1 1 1v48a.945.945 0 0 1-1 1Z"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 20.939v30.038m15.019-15.019H20.981m25.639-10.62l-21.24 21.24m21.24 0l-21.24-21.24"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M60 61H12a.945.945 0 0 1-1-1V12a.945.945 0 0 1 1-1h48a.945.945 0 0 1 1 1v48a.945.945 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}