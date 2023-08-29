package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsWithCats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.81 18.422c-.676-2.739-.57-4.269-.57-4.269M24 18.422v-4.269m-13.33.384c-.711-1.53-.433-5.397-.433-5.397c2.134 1.85 2.105 3.76 2.105 3.76l-1.672 1.637Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.602 14.003s4.803-7.614 6.973-8.503c2.277 3.415 1.853 7.7 1.853 10.788v23.474A2.738 2.738 0 0 1 37.69 42.5H10.31a2.738 2.738 0 0 1-2.738-2.738V16.288c0-3.088-.424-7.373 1.853-10.788c2.17.89 6.973 8.503 6.973 8.503h15.204Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.19 18.422c.676-2.739.57-4.269.57-4.269m9.57.384c.711-1.53.433-5.397.433-5.397c-2.134 1.85-2.105 3.76-2.105 3.76l1.672 1.637Z"/><circle cx="14.264" cy="24.119" r="2.38" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.452 42.5c0-8.883-4.301-14.23-7.1-16.768a3.507 3.507 0 0 0-4.704 0c-2.8 2.539-7.1 7.885-7.1 16.768"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.765 33.102c-2.14-.851-2.832-1.474-3.765-3.081c-.933 1.607-1.625 2.23-3.765 3.08"/><circle cx="33.736" cy="24.119" r="2.38" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}