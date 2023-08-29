package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hevy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.48 29.447V36.5a2.767 2.767 0 0 0 3.53 2.66l6.525-1.87a3.7 3.7 0 0 0 2.683-3.391c.12-2.681.282-6.884.282-10.187c0-3.133-.182-6.489-.33-8.698a3.709 3.709 0 0 0-2.519-3.27l-7.432-2.5a2.424 2.424 0 0 0-3.197 2.298v4.524a7.65 7.65 0 0 1-3.624 6.505l-4.796 2.97a7.65 7.65 0 0 0-3.624 6.505v4.047a2.767 2.767 0 0 1-3.65 2.623l-6.98-2.348a3.709 3.709 0 0 1-2.518-3.27c-.148-2.209-.33-5.565-.33-8.698c0-3.303.161-7.506.282-10.187a3.7 3.7 0 0 1 2.683-3.392l6.962-1.995a2.424 2.424 0 0 1 3.092 2.33v7.509"/>`),
		g.Group(children),
	)
}