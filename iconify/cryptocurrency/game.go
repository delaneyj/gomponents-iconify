package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Game(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-3.131-19.315v2.488h11.085v-2.488zm11.085 4.144H12.869v2.488h8.233v2.335c-1.104.952-3.266 1.444-3.266 1.444c-6.347 1.106-8.37-4.33-8.37-4.33c-1.84-6.265 2.62-8.753 2.62-8.753c5.474-3.317 9.568.921 9.568.921l1.932-1.888c-3.22-3.686-8.187-2.995-8.187-2.995c-8.05.83-9.108 7.601-9.108 7.601c-1.334 5.436 2.346 9.168 2.346 9.168c6.716 6.817 15.363.552 15.363.552V16.83z"/>`),
		g.Group(children),
	)
}