package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lovegame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.162 5.337a10.648 10.648 0 0 0-8.27 17.372h0l17.076 20.045l16.91-19.845l.084-.095l.084-.105h0A10.659 10.659 0 1 0 23.968 10a10.627 10.627 0 0 0-8.796-4.662Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.214 25.697a3.157 3.157 0 1 0 0-6.313h0a3.157 3.157 0 1 0 0-6.313m-5.241 1.105c.953-.79 1.905-1.105 3.97-1.105h1.27m-5.24 11.521c.952.79 1.746 1.105 3.97 1.105h1.27M33.596 14.63a3.894 3.894 0 0 0-3.472-1.559h-.316a4.218 4.218 0 0 0-4.261 4.209v4.209m8.523 0a4.235 4.235 0 0 1-4.262 4.209h0a4.21 4.21 0 1 1 4.262-4.21Z"/>`),
		g.Group(children),
	)
}