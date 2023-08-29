package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Halma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.28 13.12c-2 1.16-3.13 3.27-3.14 5.43c-.92.55-1.99.85-3.14.85s-2.22-.3-3.14-.85a6.242 6.242 0 0 1-3.14-5.43L24 2.5l6.28 10.62Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.86 18.55c0 1.07-.27 2.16-.84 3.15s-1.38 1.77-2.3 2.3a6.288 6.288 0 0 1-6.28 0L5.38 13.25l12.34-.13c0 2.33 1.26 4.35 3.14 5.43Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.86 29.45a6.314 6.314 0 0 1-3.14 5.43l-12.34-.13L11.44 24a6.288 6.288 0 0 0 6.28 0c.92.53 1.73 1.31 2.3 2.3s.84 2.08.84 3.15Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.28 34.88L24 45.5l-6.28-10.62c2-1.16 3.13-3.27 3.14-5.43c.92-.55 1.99-.85 3.14-.85s2.22.3 3.14.85a6.242 6.242 0 0 1 3.14 5.43Zm12.34-21.63L36.56 24a6.288 6.288 0 0 0-6.28 0c-.92-.53-1.73-1.31-2.3-2.3s-.84-2.08-.84-3.15a6.314 6.314 0 0 1 3.14-5.43l12.34.13Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.62 34.75l-12.34.13c0-2.33-1.26-4.35-3.14-5.43c0-1.07.27-2.16.84-3.15s1.38-1.77 2.3-2.3a6.288 6.288 0 0 1 6.28 0l6.06 10.75Z"/>`),
		g.Group(children),
	)
}