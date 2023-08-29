package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandslideRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-2.25l4 1.3l11.025-3.675L19.6 18.8q.75 1 .2 2.1T18 22H4q-.825 0-1.413-.588T2 20Zm4-3.05L2 15.6v-1.85l4 1.3l6.9-2.3l2.55 1.025L6 16.95Zm13.6-3.45l2.2-.975q.55-.25.875-.738T23 10.7V9.6q0-.7-.438-1.25t-1.137-.7l-2-.45q-.45-.1-.888.013t-.787.387l-1 .8q-.35.275-.55.688t-.2.862v1.1q0 .45.2.863t.55.687l.8.65q.425.35.988.413T19.6 13.5ZM6 12.95L2 11.6V10q0-.825.588-1.412T4 8h3q.475 0 .9.213t.7.587l1.975 2.625L6 12.95Zm6.975-5.35l2.775-1.1q.575-.225.913-.725T17 4.65v-2q0-.725-.45-1.288T15.4.675L12.8.15q-.4-.075-.775 0t-.725.3l-1.4.95q-.425.275-.663.725T9 3.075v1.85q0 .5.238.95T9.9 6.6l1.225.825q.425.275.913.325t.937-.15Z"/>`),
		g.Group(children),
	)
}