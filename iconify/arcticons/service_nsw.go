package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceNsw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.156 37.362c6.68-13.423 19.09-20.588 21.402-19.13c1.695 1.163 2.664 14.505-2.994 27.211"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.928 16.296c.941-1.794 1.966-3.5 3.047-3.532c1.046-.031 4.825 5.95 7.35 14.23"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.055 21.79c-.077-3.153-1.346-7.198-2.25-7.41c-1.4-.33-4.828 2.17-6.566 4.619m17.497 16.633c3.586.864 8.41 3.82 8.338 4.327a6.136 6.136 0 0 1-2.287 2.64"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.426 38.013c4.015-4.242 4.752-9.273 4.293-9.864s-4.535-.73-8.93-.114m-13.797-2.756s1.92-1.624 2.595-1.388c.693.243 3.48 10.338.499 21.04"/>`),
		g.Group(children),
	)
}