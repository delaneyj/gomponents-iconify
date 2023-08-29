package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meinheidelberg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.242 5.131C13.716 6.59 8.063 12.496 6.768 18.068A17.792 17.792 0 0 0 19.55 39.493a2.719 2.719 0 0 1 1.268.788l2.935 3.09a.453.453 0 0 0 .635 0h0l2.944-3.09a2.718 2.718 0 0 1 1.277-.779a17.812 17.812 0 0 0-9.367-34.37Zm3.479 24.279h2.718"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.874 23.521a6.259 6.259 0 0 0-2.264-.905c-1.631 0-2.9 1.812-4.53 1.812s-2.945-1.812-4.53-1.812a7.108 7.108 0 0 0-2.265.906m0-3.172v-5.435h4.53v5.436h4.53v-5.436h4.53v5.436m-13.59 9.059a2.718 2.718 0 0 1 5.436 0m2.718 0a2.718 2.718 0 0 1 5.435 0"/>`),
		g.Group(children),
	)
}