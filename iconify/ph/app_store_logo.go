package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppStoreLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m64.34 196.07l-9.45 16a8 8 0 1 1-13.78-8.14l9.46-16a8 8 0 1 1 13.77 8.14ZM232 152h-47.8l-30.73-52a8 8 0 1 0-13.77 8.14l61.41 103.93a8 8 0 0 0 13.78-8.14L193.66 168H232a8 8 0 0 0 0-16Zm-89.53 0H90.38l68.51-115.93a8 8 0 0 0-13.78-8.14L128 56.89l-17.11-29a8 8 0 1 0-13.78 8.14l21.6 36.55L71.8 152H24a8 8 0 0 0 0 16h118.47a8 8 0 1 0 0-16Z"/>`),
		g.Group(children),
	)
}