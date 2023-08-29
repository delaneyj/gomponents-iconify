package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAndroidCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 16a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm9-1a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm0 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-4a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z" clip-rule="evenodd"/><path d="m9.236 12.777l-.763-1.291l6.791-4.013l.763 1.291l-6.791 4.013Zm6.021 5.752l.772-1.286l-6.286-3.772l-.772 1.286l6.286 3.772Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}