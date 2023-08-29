package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeFundsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.378 4.514a9.962 9.962 0 0 1 6.627-2.511c5.523 0 10 4.477 10 10a9.954 9.954 0 0 1-1.793 5.715l-2.707-5.715h2.5A8 8 0 0 0 6.279 6.416l-.9-1.902Zm13.253 14.978a9.962 9.962 0 0 1-6.626 2.51c-5.523 0-10-4.476-10-10c0-2.124.663-4.094 1.793-5.714l2.707 5.715h-2.5A8 8 0 0 0 17.73 17.59l.901 1.902Zm-5.212-4.66l-2.828-2.83l-2.829 2.83l-1.414-1.415l4.243-4.242l2.828 2.828l2.828-2.829l1.415 1.415l-4.243 4.242Z"/>`),
		g.Group(children),
	)
}