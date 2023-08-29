package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarPanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 4v2.594l-1.813-1.782l-1.374 1.375L13.592 8H11v2h2.594l-1.781 1.813l1.374 1.374L15 11.408V14h2v-2.594l1.813 1.781l1.375-1.374L18.405 10H21V8h-2.594l1.782-1.813l-1.375-1.375L17 6.595V4zm1 3c1.102 0 2 .898 2 2c0 1.102-.898 2-2 2c-1.102 0-2-.898-2-2c0-1.102.898-2 2-2zm-9.781 9L4 24.875V28h24v-3.125L25.781 16zm1.562 2H24.22L26 25.125V26H6v-.875zM9 19l-.406 1.813h2.094L11 19zm4 0l-.188 1.813h2.094L15 19zm4 0l.094 1.813h2.093L19 19zm4 0l.313 1.813h2.093L23 19zM8.187 22.813L7.814 25h2.28l.313-2.188zm4.5 0L12.5 25h2.313l.093-2.188zm4.407 0L17.187 25H19.5l-.188-2.188zm4.5 0L21.906 25h2.282l-.375-2.188z"/>`),
		g.Group(children),
	)
}