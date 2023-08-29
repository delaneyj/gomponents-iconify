package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.437 8.457c-6.376.061-11.789 3.17-12.897 7.406a10.062 10.062 0 1 0 9.19 17.545l.577 5.49a.72.72 0 0 0 1.294.354l4.844-6.505a5.792 5.792 0 0 0 7.156-4.33a8.565 8.565 0 0 0 .895.05a7.747 7.747 0 0 0 8.004-7.465a7.65 7.65 0 0 0-6.846-7.383c-2.187-3.15-6.876-5.161-12.03-5.162h-.187Z"/>`),
		g.Group(children),
	)
}