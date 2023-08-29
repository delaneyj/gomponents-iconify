package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 3.5h12a.5.5 0 0 1 .5.5v.572L8 8.286L1.5 4.572V4a.5.5 0 0 1 .5-.5Zm-.5 2.8V12a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5V6.3L8.372 9.8L8 10.014L7.628 9.8L1.5 6.3ZM0 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}