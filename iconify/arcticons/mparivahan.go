package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mparivahan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.166 26.811h7.105m-7.105 1.84h7.105m-8.957 5.375v1.124a1.366 1.366 0 1 1-2.733 0v-1.124m16.275 0v1.124a1.366 1.366 0 1 1-2.733 0v-1.124m-11.426-10.64c0-1.148.957-2.059 2.906-2.059s2.332 1.039 2.332 2.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.761 31.202v-3.826c0-1.366-.305-2.522-1.275-3.06a3.177 3.177 0 0 1-1.421-1.859a168.45 168.45 0 0 0-1.603-3.753a3.757 3.757 0 0 0-2.715-1.877c-.729-.073-2.077-.145-5.028-.145s-4.3.073-5.029.145a3.757 3.757 0 0 0-2.714 1.877a165.73 165.73 0 0 0-1.603 3.753a3.178 3.178 0 0 1-1.422 1.858c-.97.54-1.275 1.695-1.275 3.061v3.826M38.5 23.386H18.937"/><circle cx="36.935" cy="26.811" r="1.366" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="20.502" cy="26.811" r="1.366" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.719 34.026h10.657a1.621 1.621 0 0 0 1.622-1.622h0a1.621 1.621 0 0 0-1.621-1.621H18.06a1.621 1.621 0 0 0-1.622 1.621h0a1.621 1.621 0 0 0 1.622 1.622Zm6.021-10.64c0-1.148-.957-2.059-2.906-2.059s-2.332 1.039-2.332 2.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.719 16.682V7.099a2.6 2.6 0 0 0-2.6-2.599H9.601a2.6 2.6 0 0 0-2.599 2.6v33.8a2.6 2.6 0 0 0 2.6 2.6h16.517a2.6 2.6 0 0 0 2.6-2.6v-6.874"/>`),
		g.Group(children),
	)
}