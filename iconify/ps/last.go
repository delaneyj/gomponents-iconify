package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Last(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M370 329q-63-1-91-66l-5-11l-44-101q-11-26-35.5-42.5T139 92q-41 0-70.5 29.5T39 192t29.5 70.5T139 292q61 0 88-53l18 41q-40 50-106 50q-57 0-97-40.5T2 192t40-97.5T139 54q89 0 127 85q1 2 7.5 17.5t18 42T311 244q7 16 12.5 25t17 16t27.5 7q25 1 41-12t16-33q0-9-2.5-15.5t-10-11.5t-14-7.5T378 205q-44-15-63-31.5T296 126q0-32 20-50.5T370 57q45 0 67 40l-29 15q-16-22-39-22q-16 0-26.5 10T332 125q0 6 1 11t5 9t6.5 6t10 5.5T365 161l13.5 4.5L393 170q37 12 53 29t16 49q0 35-26 58t-66 23z"/>`),
		g.Group(children),
	)
}