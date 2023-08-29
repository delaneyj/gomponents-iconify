package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Algo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0c8.837 0 16 7.163 16 16s-7.163 16-16 16S0 24.837 0 16S7.163 0 16 0zm3.093 7h-2.4l-.056.084l-2.246 3.888l-2.302 3.986l-2.287 3.972L7.5 22.916h2.75l2.303-3.986l2.301-3.972l2.288-3.986l.38-.632l.167.632l.702 2.624l-.786 1.362l-2.301 3.972l-2.288 3.986h2.75l2.302-3.986l1.193-2.063l.562 2.063l1.066 3.986h2.47l-1.066-3.986l-1.067-3.972l-.28-1.025l1.712-2.961H20.16l-.085-.295l-.87-3.256L19.093 7z"/>`),
		g.Group(children),
	)
}