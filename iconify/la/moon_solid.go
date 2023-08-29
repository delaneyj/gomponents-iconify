package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v2H3v2h2v2h2V9h2V7H7V5zm15.281 2.938L18.625 8C13.281 8.191 9 12.578 9 17.969c0 5.511 4.488 10 10 10c5.39 0 9.777-4.282 9.969-9.625l.062-1.625l-1.468.687A5.94 5.94 0 0 1 25 17.97c-3.324 0-6-2.676-6-6c0-.914.223-1.75.594-2.531zm-2.906 2.375c-.125.554-.375 1.062-.375 1.656c0 4.406 3.594 8 8 8c.605 0 1.121-.246 1.688-.375c-.762 3.625-3.829 6.375-7.688 6.375c-4.43 0-8-3.57-8-8c0-3.852 2.758-6.887 6.375-7.657z"/>`),
		g.Group(children),
	)
}