package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnyunContinuingEducation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.782 5.5c3.376 11.031 9.66 21.056 17.744 30.458c-3.047-3.233-4.064-7.407-13.665-5.37a40.151 40.151 0 0 1-4.419-9.383a60.145 60.145 0 0 1-5.439 12.986c-3.478 2.559-8.449 2.465-13.53 2.175c7.81-8.652 15.247-17.75 19.308-30.865Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.474 36.366c1.605 2.36 4.928 6.417 9.96 6.118c16.068-.952 14.87-9.978 27.092-6.526M19.003 34.19a33.61 33.61 0 0 1 9.857-3.603"/>`),
		g.Group(children),
	)
}