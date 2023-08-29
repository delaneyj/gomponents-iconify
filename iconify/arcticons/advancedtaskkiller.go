package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Advancedtaskkiller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.15 20.67a3.07 3.07 0 1 1 3.07 3.07h-.01a3.07 3.07 0 0 1-3.06-3.07Zm18.64 3.07a3.07 3.07 0 1 1 3.07-3.07h0a3.07 3.07 0 0 1-3.07 3.07Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 42h33.068a2 2 0 0 0 2-2.001c-.003-4.31-.012-14.183-.018-15.489A18.535 18.535 0 0 0 39 13.72l2.58-2.56a3.016 3.016 0 1 0-4.27-4.26h0l-2.55 2.6a18.39 18.39 0 0 0-21.51 0l-2.57-2.57A3.22 3.22 0 0 0 8.5 6a3.002 3.002 0 0 0-2.987 3.506c.108.662.471 1.257.95 1.725L9 13.71a18.487 18.487 0 0 0-3.47 10.8c-.002 3.312-.034 11.64-.05 15.486a1.996 1.996 0 0 0 2 2.004Zm11.65-21.28a2.93 2.93 0 1 1-2.94-2.92h.01a2.93 2.93 0 0 1 2.94 2.9l-.01.02Zm12.65 2.92a2.92 2.92 0 1 1 2.92-2.92h0a2.92 2.92 0 0 1-2.91 2.9l-.01.02ZM19.5 36.967l9-9m-9 0l9 9"/>`),
		g.Group(children),
	)
}