package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nagad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.88 6.397a19.347 19.347 0 1 0 23.516 12.988"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.92 25.893c-.062-4.213 2.564-12.09 11.48-15.938l-3.603-6.412c-5.374 2.626-12.58 11.664-7.878 22.35Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.136 12.409a16.778 16.778 0 0 0-2.921 8.842c1.839-3.791 7.725-9.648 17.419-9.077l-.336-7.347a17.994 17.994 0 0 0-10.613 3.856"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.465 12.153a16.506 16.506 0 0 0-10.569 8.127c3.347-2.56 11.238-5.145 19.64-.278l3.002-6.714a17.746 17.746 0 0 0-9.924-1.532"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.438 20.07a11.604 11.604 0 1 0 19.046-2.772"/>`),
		g.Group(children),
	)
}