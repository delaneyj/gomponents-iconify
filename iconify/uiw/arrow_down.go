package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.364.765c-.037 5.419-.078 10.936-.12 16.553a398914.9 398914.9 0 0 0-6.454-6.361c-.385-.312-.802-.221-1.067.035c-.248.24-.353.7.007 1.09c2.368 2.326 4.975 4.892 7.82 7.697a.794.794 0 0 0 .554.222a.745.745 0 0 0 .539-.222c2.726-2.75 5.287-5.335 7.683-7.755a.754.754 0 0 0-.055-1.032c-.371-.374-.885-.229-1.093 0a3545.93 3545.93 0 0 1-6.386 6.437c.041-5.609.08-11.163.117-16.664c0-.26-.212-.765-.767-.765s-.778.469-.778.765Z"/>`),
		g.Group(children),
	)
}