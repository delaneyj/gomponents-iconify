package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeCnyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.378 4.514a9.962 9.962 0 0 1 6.627-2.511c5.523 0 10 4.477 10 10a9.954 9.954 0 0 1-1.793 5.715l-2.707-5.715h2.5A8 8 0 0 0 6.279 6.416l-.9-1.902Zm13.253 14.978a9.962 9.962 0 0 1-6.626 2.51c-5.523 0-10-4.476-10-10c0-2.124.663-4.094 1.793-5.714l2.707 5.715h-2.5A8 8 0 0 0 17.73 17.59l.901 1.902Zm-5.626-5.954h3v2h-3v2h-2v-2h-3v-2h3v-1h-3v-2h2.586l-2.122-2.12l1.415-1.415l2.12 2.121l2.122-2.121l1.414 1.414l-2.12 2.122h2.585v2h-3v1Z"/>`),
		g.Group(children),
	)
}