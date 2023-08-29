package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PizzaBoyGbc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="22.245" cy="23.357" r="3.617" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.469" cy="23.808" r="3.979" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="27.059" cy="34.697" r="2.962" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.99 13.27c-.33 1.75-.91 4.02-1.65 6.48"/><ellipse cx="42.57" cy="28.808" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.537" ry="1.834" transform="rotate(-16.754 42.57 28.807)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.7 19.844c-3.32-3.19-6.99-5.48-8.98-6.24a2.79 2.79 0 0 0-.34-.12a2.542 2.542 0 0 1-1.31-3.34c.56-1.29 2.01-1.68 3.34-1.32c1.91.51 6.05 3.01 9.96 6.55c3.92 3.53 7.61 8.09 8.63 12.7m-4.86 1.46c-.55-1.89-1.52-3.7-2.71-5.39m-11.51 13.3c-4.09.96-8.4 1.6-12.4 1.86c2.02-1.93 4.54-7.33 6.63-13m9.26 10.2c2.04-.61 3.98-1.3 5.73-2.07c.72-.32.24 2.24 2.2 2.03c.99-.1 1.52-.79 1.66-2.66c.11-1.46 1.79-2.69 2.34-3.2"/><circle cx="26.362" cy="18.615" r=".75" fill="currentColor"/><circle cx="25.862" cy="29.308" r=".75" fill="currentColor"/><circle cx="19.681" cy="36.244" r=".75" fill="currentColor"/><circle cx="35.638" cy="30.71" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.26 35.27l-3.2.9c-1.04.29-2.12-.32-2.41-1.36L5.07 14.92a1.96 1.96 0 0 1 1.35-2.41l11.64-3.27c1.04-.29 2.11.32 2.41 1.36l2.23 7.93"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.88 20.644l-9.25 2.59a.525.525 0 0 1-.64-.36l-2.17-7.72c-.07-.27.09-.56.36-.63l9.42-2.65c.28-.07.56.09.64.36l2.16 7.72c.03.11.03.22-.02.31m-7.603 6.059l.931 3.319m1.194-2.125l-3.319.931"/><circle cx="19.175" cy="26.722" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}