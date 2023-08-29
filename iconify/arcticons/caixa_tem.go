package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaixaTem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.387C13.23 7.387 4.5 14.825 4.5 24a14.827 14.827 0 0 0 3.147 9.045L4.5 40.613l9.116-2.554A21.941 21.941 0 0 0 24 40.613c10.77 0 19.5-7.438 19.5-16.613S34.77 7.387 24 7.387Zm1.988 11.312h6.342"/><ellipse cx="18.273" cy="18.699" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.029" ry="3.408"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 34.839c5.064 0 10.495-4.307 10.495-10.082c-4.023 4.07-7.087 4.686-10.495 4.686s-6.532-.615-10.555-4.686c0 5.775 5.49 10.082 10.555 10.082Z"/>`),
		g.Group(children),
	)
}