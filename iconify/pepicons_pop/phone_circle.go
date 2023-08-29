package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="m15.52 13.192l-.581.582a31.121 31.121 0 0 1-2.827-2.83l.58-.58a3 3 0 0 0 0-4.243l-1.414-1.414a3 3 0 0 0-4.242 0L4.23 7.513A1 1 0 0 0 4 8.57c2.374 6.36 6.818 10.834 13.308 13.312a1 1 0 0 0 1.064-.227l2.806-2.806a3 3 0 0 0 0-4.242l-1.415-1.415a3 3 0 0 0-4.242 0Zm4.244 2.829a1 1 0 0 1 0 1.414l-2.342 2.341c-5.424-2.23-9.173-6-11.317-11.31L8.45 6.12a1 1 0 0 1 1.414 0l1.414 1.415a1 1 0 0 1 0 1.414l-1.234 1.234a1 1 0 0 0-.063 1.345a33.163 33.163 0 0 0 4.373 4.376a1 1 0 0 0 1.345-.063l1.236-1.235a1 1 0 0 1 1.414 0l1.415 1.414Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}