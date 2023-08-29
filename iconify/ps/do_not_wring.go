package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotWring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m260 149l30-38q-22-4-34-4q-18 0-34 6l30 36h8zm-4 207l-21 26q6 2 21 2t21-2zM73 380l100-124L51 107H0v277h43q8 0 30-4zM43 149q22 0 40 13t26 34l17 49l-17 49q-8 21-26 34t-40 13V149zm416-40L339 256l103 124q16 4 29 4h43V107h-43q-8 0-12 2zm10 232q-22 0-40-13t-26-34l-17-49l17-49q8-21 26-34t40-13v192zM356 166L465 34q5-6 4.5-15.5T463 4q-8-4-15-4q-12 0-17 9L329 134l-28 35l-45 53l-45-56l-28-34L81 9q-3-5-4-5q-1-1-9-4h-4q-9 0-13 4q-6 5-7 14.5T49 34l109 132l30 39l42 51l-38 45l-30 34L47 478q-11 15 2 30q8 4 15 4q12 0 17-9l113-138l28-34l34-43l34 43l28 34l113 138q7 9 17 9q9 0 13-4q6-5 7-14.5t-5-15.5L348 335l-28-34l-36-45l42-51z"/>`),
		g.Group(children),
	)
}