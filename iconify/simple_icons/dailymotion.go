package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dailymotion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.068 11.313a3.104 3.104 0 0 0-3.104 3.11c0 1.753 1.35 3.085 3.255 3.085l-.016.002c1.59 0 2.925-1.31 2.925-3.04c0-1.8-1.336-3.157-3.062-3.157zM0 0v24h24V0H0zm20.693 20.807h-3.576v-1.41c-1.1 1.08-2.223 1.47-3.715 1.47c-1.522 0-2.832-.495-3.93-1.485c-1.448-1.275-2.198-2.97-2.198-4.936c0-1.8.7-3.414 2.01-4.674c1.17-1.146 2.595-1.73 4.185-1.73c1.52 0 2.69.513 3.53 1.59V4.157l3.693-.765V3.39l.002.003h-.002v17.414z"/>`),
		g.Group(children),
	)
}