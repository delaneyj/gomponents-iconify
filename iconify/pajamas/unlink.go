package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 1.75a.75.75 0 0 0-1.5 0V4.5H1.75a.75.75 0 0 0 0 1.5H6V1.75Zm5.151 1.16a1.371 1.371 0 0 1 1.94 1.939l-2.871 2.87a.75.75 0 1 0 1.06 1.061l2.871-2.87a2.871 2.871 0 0 0-4.06-4.061L7.22 4.719A.75.75 0 0 0 8.28 5.78l2.871-2.87ZM10 14.25a.75.75 0 0 0 1.5 0V11.5h2.75a.75.75 0 0 0 0-1.5H10v4.25ZM5.781 7.22a.75.75 0 0 1 0 1.06l-2.87 2.871a1.371 1.371 0 0 0 1.939 1.94l2.87-2.871a.75.75 0 1 1 1.061 1.06l-2.87 2.871a2.871 2.871 0 0 1-4.061-4.06L4.72 7.22a.75.75 0 0 1 1.061 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}