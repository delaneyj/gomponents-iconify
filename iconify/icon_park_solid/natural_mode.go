package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NaturalMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNaturalMode0"><g fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="4"><path d="M24 27c-5.657 0-9.935-4.343-9.935-10L14 7l5 4l5-6l5 6l5-4v10c0 5.657-4.342 10-10 10Zm-.514 16.314c1.562-1.562-.337-5.995-4.242-9.9c-3.906-3.905-8.338-5.805-9.9-4.242c-1.562 1.562.337 5.994 4.243 9.9c3.905 3.904 8.337 5.804 9.9 4.242Z"/><path d="M24.829 42.97c1.562 1.563 5.994-.337 9.9-4.242c3.905-3.905 5.804-8.338 4.242-9.9c-1.562-1.562-5.994.338-9.9 4.243c-3.905 3.905-5.804 8.337-4.242 9.9Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNaturalMode0)"/>`),
		g.Group(children),
	)
}