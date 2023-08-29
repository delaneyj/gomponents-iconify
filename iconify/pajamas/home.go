package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.38 1.353L8 1.131l-.38.222l-7.25 4.25a.75.75 0 0 0 .76 1.294l.87-.51V14h12V6.387l.87.51a.75.75 0 1 0 .76-1.294l-7.25-4.25Zm4.12 4.154L8 2.87L3.5 5.507V12.5H6V8h4v4.5h2.5V5.507ZM8.5 9.5v3h-1v-3h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}