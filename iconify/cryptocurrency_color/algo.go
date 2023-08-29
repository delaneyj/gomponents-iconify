package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Algo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#000"/><path fill="#FFF" d="m10.25 22.916l2.303-3.986l2.301-3.972l2.288-3.986l.38-.632l.167.632l.702 2.624l-.786 1.362l-2.301 3.972l-2.288 3.986h2.75l2.302-3.986l1.193-2.063l.562 2.063l1.066 3.986h2.47l-1.066-3.986l-1.067-3.972l-.28-1.025l1.712-2.961H20.16l-.085-.295l-.87-3.256L19.093 7h-2.4l-.056.084l-2.246 3.888l-2.302 3.986l-2.287 3.972L7.5 22.916z"/></g>`),
		g.Group(children),
	)
}