package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Handdrag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M892 277q-31 80-43 173q-16 126-16 190v32q0 45-13.5 97.5t-32 96T750 946t-32 57l-13 21H321q-6-5-14.5-14T276 975.5T237 927t-30.5-50t-13.5-45q0-38-43.5-122T62 561l-43-65Q0 477 0 449.5t19-47T64.5 383t45.5 20q3 3 7 9.5t16.5 24.5t24.5 34t28 36t28.5 34.5t24 24.5t18.5 10q-2-32 0-32l-64-405q-4-26 10.5-48T242 65t43.5 11.5T309 119l40 237q36 113 36 28V64q0-26 18.5-45t45-19T494 19t19 45v384q7-2 14.5-10.5T538 416l50-297q4-27 24-42.5T656 65t38 26t10 48q-3 16-8.5 41.5t-14 90.5t-8.5 113q0 32 8 64t16 48l8 16q4-10 12-22l64-257q9-25 32-36t46-2t33 33t0 49z"/>`),
		g.Group(children),
	)
}