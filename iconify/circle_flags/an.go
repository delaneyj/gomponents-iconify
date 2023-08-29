package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func An(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsAn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAn0)"><path fill="#eee" d="M0 0h171l85 32l85-32h171v171l-32 85l32 85v171H341l-85-32l-85 32H0V341l32-85l-32-85Z"/><path fill="#d80027" d="M171 0h170v512H171z"/><path fill="#0052b4" d="M512 171v170H0V171z"/><path fill="#eee" d="m236 247l52-37h-64l52 37l-20-61zm-45 79l52-37h-64l52 37l-20-61zm90 0l52-37h-64l52 37l-20-61zm74-47l52-37h-64l52 37l-20-61zm-238 0l52-37h-64l52 37l-20-61z"/></g>`),
		g.Group(children),
	)
}