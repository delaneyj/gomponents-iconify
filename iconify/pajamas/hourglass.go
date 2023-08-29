package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.75 0a.75.75 0 0 0 0 1.5H3v.593c0 1.26.5 2.468 1.391 3.359L6.94 8l-2.548 2.548A4.75 4.75 0 0 0 3 13.907v.593h-.25a.75.75 0 0 0 0 1.5h10.5a.75.75 0 0 0 0-1.5H13v-.593c0-1.26-.5-2.468-1.391-3.359L9.06 8l2.548-2.548A4.75 4.75 0 0 0 13 2.093V1.5h.25a.75.75 0 0 0 0-1.5H2.75Zm8.75 1.5h-7v.593c0 .69.219 1.356.618 1.907h5.764a3.25 3.25 0 0 0 .618-1.907V1.5ZM8 6.94L6.56 5.5h2.88L8 6.94Zm3.5 7.56v-.593a3.25 3.25 0 0 0-.952-2.298L8 9.06l-2.548 2.548a3.25 3.25 0 0 0-.952 2.298v.593h7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}