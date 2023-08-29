package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoronoiMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M2.5 12.5A2.5 2.5 0 0 0 0 15v70a2.5 2.5 0 0 0 2.5 2.5h95A2.5 2.5 0 0 0 100 85V15a2.5 2.5 0 0 0-2.5-2.5zm2.5 5h43.745l-.033 5.709l-16.68 27.883L5 64.367Zm46.747 0H95v26.108l-25.982.397l-17.304-20.853Zm-1.432 8.857l16.36 19.53l-5.86 12.244l-9.934 1.396l-15.722-7.808zM95 46.761V82.5h-7.232L63.637 59.213l5.84-12.204Zm-61.991 7.184l16.15 8.053L50.966 82.5H5V67.674Zm28.26 7.15L83.448 82.5H53.98L52.2 62.371Z" color="currentColor"/><path fill="currentColor" d="M60.366 68.73a4 4 0 0 0-4 4a4 4 0 0 0 4 4a4 4 0 0 0 4-4a4 4 0 0 0-4-4zm19.096-13.746a4 4 0 0 0-4 4a4 4 0 0 0 4 4a4 4 0 0 0 4-4a4 4 0 0 0-4-4zm-4.901-28.179a4 4 0 0 0-4 4a4 4 0 0 0 4 4a4 4 0 0 0 4-4a4 4 0 0 0-4-4zm-50.582-.7a4 4 0 0 0-4 4a4 4 0 0 0 4 4a4 4 0 0 0 4-4a4 4 0 0 0-4-4zm29.929 16.802a4 4 0 0 0-4 4a4 4 0 0 0 4 4a4 4 0 0 0 4-4a4 4 0 0 0-4-4zM39.03 64.435a4 4 0 0 0-4 4a4 4 0 0 0 4 4a4 4 0 0 0 4-4a4 4 0 0 0-4-4z" color="currentColor"/>`),
		g.Group(children),
	)
}