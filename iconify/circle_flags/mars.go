package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsMars0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMars0)"><path fill="#6da544" d="M170 0h172l32 256l-32 256H170l-32-256Z"/><path fill="#d80027" d="M0 0h170v512H0Z"/><path fill="#0052b4" d="M342 0h170v512H342Z"/></g>`),
		g.Group(children),
	)
}