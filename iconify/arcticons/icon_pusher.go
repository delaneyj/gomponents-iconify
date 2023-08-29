package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IconPusher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.46 27.895c.241 1.105 1.104 5.97.587 6.592s-4.179 3.847-4.694 4.658a1.353 1.353 0 0 0-.241.655c.103 1.192 2.346 1.192 2.346 2.14c0 .725-2.933.619-4.934-.06c-.84-.338-.242-1.803.686-2.421l3.196-5.183a2.346 2.346 0 0 0-.76-1.36a26.76 26.76 0 0 1-4.848 3.676a34.202 34.202 0 0 1-4.055 3.347c.19.932 2.191 1.242 2.467 2.019s-2.001.495-2.57.48a14.23 14.23 0 0 1-2.727-.843c-.323-.118-.009-1.417 1.122-2.347c0 0 3.192-3.554 3.796-4.082a34.548 34.548 0 0 0 3.537-5.01a4.978 4.978 0 0 1 2.214-4.55c.173-1.45 2.968-5.36 3.773-6.218c.334-.355.018-2.243 3.313-2.709s7.702-1.99 8.754-2.749"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.46 27.895a11.144 11.144 0 0 0 2.07-3.071c1.364-2.882 1.83-3.296 1.778-5.297a24.144 24.144 0 0 0 7.207-4.159c.115-.908.448-4.451.448-4.451m-9.294 5.763c-.265-.99 1.148-2.916 3.094-2.646a1.515 1.515 0 0 1 1.46 1.59"/>`),
		g.Group(children),
	)
}