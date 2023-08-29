package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColdChain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M30 4c-2.761 0-5 2.398-5 5.357v20.398a8 8 0 1 0 10 0V9.357C35 6.398 32.761 4 30 4Zm3 18.061c-1.457.024-2.822-.33-4.018-.64c-.73-.19-1.396-.362-1.982-.422V9.357c0-1.775 1.343-3.214 3-3.214s3 1.439 3 3.214v.619h-2a1 1 0 0 0 0 2h2v2h-2a1 1 0 1 0 0 2h2v2h-2a1 1 0 1 0 0 2h2v2.085Z" clip-rule="evenodd"/><path d="M16.33 12a1 1 0 1 0-2 0v3.268l-2.83-1.634a1 1 0 1 0-1 1.732L13.33 17l-2.83 1.634a1 1 0 1 0 1 1.732l2.83-1.634V22a1 1 0 0 0 2 0v-3.268l2.83 1.634a1 1 0 1 0 1-1.732L17.33 17l2.83-1.634a1 1 0 1 0-1-1.732l-2.83 1.634V12Z"/></g>`),
		g.Group(children),
	)
}