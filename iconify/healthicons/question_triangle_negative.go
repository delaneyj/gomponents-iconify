package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionTriangleNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsQuestionTriangleNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24.9 6.849a1 1 0 0 0-1.8 0L6.7 40.563A1 1 0 0 0 7.598 42H40.4a1 1 0 0 0 .9-1.437L24.899 6.849Zm4.598 18.412c0-1.486-.777-2.595-1.807-3.28c-.982-.653-2.201-.94-3.344-.977c-1.151-.037-2.374.176-3.414.68c-1.038.503-2.04 1.38-2.385 2.721a1.5 1.5 0 0 0 2.905.746c.056-.215.252-.507.789-.767c.534-.259 1.26-.405 2.01-.381c.756.023 1.386.216 1.778.476c.345.23.468.47.468.782c0 .373-.07.588-.13.707a.698.698 0 0 1-.24.272c-.254.175-.674.285-1.228.322a1.5 1.5 0 0 0-1.4 1.497V30.5a1.5 1.5 0 0 0 3 0v-1.176a4.522 4.522 0 0 0 1.33-.614c1.03-.71 1.668-1.87 1.668-3.449ZM27 36a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsQuestionTriangleNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}