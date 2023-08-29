package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m12.52 10.192l-.581.582a31.121 31.121 0 0 1-2.827-2.83l.58-.58a3 3 0 0 0 0-4.243L8.278 1.707a3 3 0 0 0-4.242 0L1.23 4.513A1 1 0 0 0 1 5.57c2.374 6.36 6.818 10.834 13.308 13.312a1 1 0 0 0 1.064-.227l2.806-2.806a3 3 0 0 0 0-4.242l-1.415-1.415a3 3 0 0 0-4.242 0Zm4.244 2.829a1 1 0 0 1 0 1.414l-2.342 2.341c-5.424-2.23-9.173-6-11.317-11.31L5.45 3.12a1 1 0 0 1 1.414 0l1.414 1.415a1 1 0 0 1 0 1.414L7.044 7.184a1 1 0 0 0-.063 1.345a33.163 33.163 0 0 0 4.373 4.376a1 1 0 0 0 1.345-.063l1.236-1.235a1 1 0 0 1 1.414 0l1.415 1.414Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}