package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriggerSource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.25.75a.75.75 0 0 1 1.5 0V2c0 2.9 2.35 5.25 5.25 5.25h1.25a.75.75 0 0 1 0 1.5H14A5.25 5.25 0 0 0 8.75 14v1.25a.75.75 0 0 1-1.5 0V14c0-2.9-2.35-5.25-5.25-5.25H.75a.75.75 0 0 1 0-1.5H2C4.9 7.25 7.25 4.9 7.25 2V.75ZM5.095 8A6.78 6.78 0 0 1 8 10.905A6.78 6.78 0 0 1 10.905 8A6.78 6.78 0 0 1 8 5.095A6.78 6.78 0 0 1 5.095 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}