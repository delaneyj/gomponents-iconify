package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IExamQualification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M28.753 6.341A1 1 0 0 0 27 7v7a2 2 0 0 0 2 2h6a1 1 0 0 0 .753-1.659l-7-8ZM20.75 23h-1.5l.75-1.8l.75 1.8Zm6.808-18L37 15.387V40a3 3 0 0 1-3 3H14a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h13.558Zm-5.712 10.23a2 2 0 0 0-3.692 0l-5 12a2 2 0 0 0 3.692 1.54l.737-1.77h4.834l.737 1.77a2 2 0 0 0 3.692-1.54l-.103-.246c.084.01.17.016.257.016h1v1a2 2 0 1 0 4 0v-1h1a2 2 0 1 0 0-4h-1v-1a2 2 0 1 0-4 0v1h-1c-.648 0-1.224.308-1.59.786l-3.564-8.555ZM15 31a2 2 0 1 0 0 4a2 2 0 1 0 0 4h12a2 2 0 0 0 .002-4H33a2 2 0 1 0 0-4H15Z"/>`),
		g.Group(children),
	)
}