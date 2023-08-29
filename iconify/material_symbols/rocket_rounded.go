package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RocketRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.725 20.525l-2.35.925q-.5.2-.938-.1T4 20.525v-4.45q0-.5.238-.95T4.9 14.4l1.1-.725q.175 2.025.537 3.5t1.188 3.35Zm3.5-17.3q.15-.15.363-.212T12 2.95q.2 0 .413.063t.362.212Q14.5 4.8 15.5 7.337t1 5.338q0 1.925-.325 3.375t-1 3.175q-.125.325-.375.55t-.6.225H9.8q-.35 0-.6-.225t-.375-.55q-.675-1.725-1-3.188T7.5 12.675q0-2.8 1-5.338t2.725-4.112ZM12 13q.825 0 1.413-.587T14 11q0-.825-.588-1.413T12 9q-.825 0-1.413.588T10 11q0 .825.588 1.413T12 13Zm4.275 7.525q.825-1.875 1.188-3.35t.537-3.5l1.1.725q.425.275.663.725t.237.95v4.45q0 .525-.437.825t-.938.1l-2.35-.925Z"/>`),
		g.Group(children),
	)
}