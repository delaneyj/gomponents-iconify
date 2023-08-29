package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CatsAndSoup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.67 15.282c2.024.716 2.66 2.32 2.66 2.32m-.348 20.835c2.437-1.489 6.164 1.666 3.786 3.734c-2.471 2.149-6.124 1.773-6.34-1.72c-.214-3.492.592-3.68 2.123-4.324c-.752-3.492.778-11.286.778-11.286m9.805.881c.73.165 3.272.234 3.492 0m-8.283-7.577c-.03.931 0 3.566 0 3.566m4.403-.776c.896 1.062 1.43 1.006 2.508 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.623 25.865c4.321-2.751 1.037-8.436.43-10.153c3.008-4.03 2.73-10.951 1.235-11.208c-.879-.151-5.55 4.01-7.199 5.889c0 0-2.31-.645-4.089-.645c-1.779 0-4.089.645-4.089.645c-1.386-1.797-6.268-5.877-7.199-5.889c-1.517-.02-1.773 7.178 1.236 11.208c-.058 2.445-2.892 8.189.43 10.153"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.33 15.282c-2.024.716-2.66 2.32-2.66 2.32m.347 20.835c-2.436-1.489-6.163 1.666-3.785 3.734c2.471 2.149 6.124 1.773 6.34-1.72c.214-3.492-.592-3.68-2.123-4.324c.752-3.492-.653-11.286-.653-11.286m-5.139-6.696c.03.931 0 3.566 0 3.566"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.841 42.83c2.264.093 11.803-.089 16.318 0M22.525 9.86c.744-.89 2.651-1.777 2.651-1.777"/>`),
		g.Group(children),
	)
}