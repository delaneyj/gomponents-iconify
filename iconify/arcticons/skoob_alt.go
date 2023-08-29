package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkoobAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="30.934" cy="18.966" r="5.244" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.934" cy="21.123" r="1.583" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.084" cy="18.966" r="5.244" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="17.084" cy="21.123" r="1.583" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.009 21.618a1.243 1.243 0 0 0-1.187 1.286a8.602 8.602 0 0 0 1.187 2.473a8.531 8.531 0 0 0 1.187-2.473a1.238 1.238 0 0 0-1.187-1.286ZM12.523 33.282c-.92-.481-1.77-1.032-1.77-2.07v-2.573l11.2 5.799a8.115 8.115 0 0 1 2.054 2.017a8.106 8.106 0 0 1 2.054-2.017l11.199-5.799v2.572c0 1.04-.849 1.59-1.769 2.071l-11.486 6.01Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.765 41.266a8.938 8.938 0 0 1-5.498-1.565a4.946 4.946 0 0 1-2.031-3.68c-1.397-.64-3.862-1.768-3.862-5.432v-5.802l3.859 2.077v-.99a8.505 8.505 0 0 1-.99-13.005s-1.385-1.401-1.385-2.23V4.5l11.565 5.777C23.983 11.555 24 12.914 24 12.914s.017-1.359 2.577-2.637L38.143 4.5v6.139c0 .828-1.385 2.23-1.385 2.23a8.505 8.505 0 0 1-.99 13.006v.99l3.858-2.078v5.802c0 3.664-2.464 4.792-3.861 5.431a4.944 4.944 0 0 1-2.032 3.68a8.937 8.937 0 0 1-5.497 1.566c-2.256-.005-4.242 2.234-4.242 2.234s-1.986-2.239-4.241-2.234Z"/>`),
		g.Group(children),
	)
}