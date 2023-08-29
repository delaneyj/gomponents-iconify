package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Networksurvey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.056 7.182v29.843m12.581-24.747v29.843M42.5 35.722L30.637 42.12l-12.58-5.096L5.5 41.335V11.492l12.556-4.31l12.58 5.095L42.5 5.88Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.677 28.775s3.103-9.466 3.103-13.531"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.707 26.366a26.032 26.032 0 0 1-2.928-11.121"/><ellipse cx="11.779" cy="15.277" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".828" ry=".633" transform="rotate(-48.828 11.78 15.277)"/><path fill="none" stroke="currentColor" stroke-dasharray="3 3" stroke-linecap="round" stroke-linejoin="round" d="M9.5 35.81s2.066-11.213 15.58-7.697c15.612 4.063 13.876-8.494 13.25-12.766"/><path fill="none" stroke="currentColor" d="m9.614 26.35l4.27-1.719m-3.537-.759l2.58-1.038"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.04 18.385a2.293 2.293 0 0 1-1.596-2.324a4.384 4.384 0 0 1 1.596-3.203m3.24 4.634a4.384 4.384 0 0 0 1.595-3.203a2.292 2.292 0 0 0-1.595-2.323"/>`),
		g.Group(children),
	)
}