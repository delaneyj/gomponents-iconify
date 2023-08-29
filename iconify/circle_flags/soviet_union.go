package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SovietUnion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsSovietUnion0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSovietUnion0)"><path fill="#d80027" d="M0 0h512v512H0z"/><path fill="#ffda44" d="m90 243l25 25l21-21l115 115l18-18l-115-115l29-29l-44-6z"/><path fill="#ffda44" d="M173 166a93 93 0 0 1 44 40c22 38 16 83-14 100c-22 13-52 7-75-15l-58 53l18 18l42-46c27 24 63 30 90 14c36-21 45-73 20-116a95 95 0 0 0-67-48zm-43-6l105-77H105l105 77l-40-124z"/><path fill="#d80027" d="m153 129l45-34h-56l45 34l-17-54z"/></g>`),
		g.Group(children),
	)
}