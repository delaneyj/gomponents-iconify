package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Page(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M80.277 12.498h-.005v-4.02a1.73 1.73 0 0 0-1.73-1.73h-32.87l-25.95 25.95v58.819c0 .956.774 1.73 1.73 1.73h57.09a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005V12.498zM29.679 83.292V36.158h17.723a1.73 1.73 0 0 0 1.73-1.73V16.705H70.32v66.587H29.679z"/>`),
		g.Group(children),
	)
}