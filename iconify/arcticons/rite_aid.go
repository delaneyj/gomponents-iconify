package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RiteAid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.39 8.003c-1.144.666-1.921 1.89-1.921 3.309s.777 2.643 1.92 3.308c1.144-.665 1.921-1.89 1.921-3.308s-.777-2.643-1.92-3.309Zm6.334 6.747c-1.28-.338-2.695-.021-3.698.982s-1.319 2.418-.98 3.697a3.825 3.825 0 0 0 3.697-.981a3.825 3.825 0 0 0 .98-3.698Zm-12.669 0c1.279-.338 2.695-.021 3.697.982s1.32 2.418.981 3.697a3.825 3.825 0 0 1-3.697-.981a3.825 3.825 0 0 1-.981-3.698Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c10.707-2.604 16.272-11.687 16.272-16.628V4.5H7.728v22.372c0 4.941 5.565 14.024 16.272 16.628Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 38.224h3.851c3.656-1.302 8.665-7.162 8.665-10.217v-5.109H11.484v5.109c0 3.055 5.009 8.915 8.665 10.217H24Zm7.943-2.754H16.057m8.55-15.126c2.804-4.358 6.06-5.46 6.06-5.46c.5-2.003 2.354-3.956 2.354-3.956l3.456 3.105c-1.553 1.703-3.105 1.953-3.105 1.953c-.501 2.605-2.555 4.358-2.555 4.358h-6.21Z"/>`),
		g.Group(children),
	)
}