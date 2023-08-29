package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunningShirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.25 31H6a2 2 0 0 1-2-2V13.75A1.752 1.752 0 0 1 5.75 12H6a1 1 0 0 0 1-1V2.75A1.752 1.752 0 0 1 8.75 1h3.5A1.752 1.752 0 0 1 14 2.75V7a2 2 0 0 0 .09.593A4.73 4.73 0 0 0 16 8a4.73 4.73 0 0 0 1.91-.407A2 2 0 0 0 18 7V2.75A1.752 1.752 0 0 1 19.75 1h3.5A1.752 1.752 0 0 1 25 2.75l.025 8.477A1 1 0 0 0 26 12h.25A1.752 1.752 0 0 1 28 13.75v15.5A1.752 1.752 0 0 1 26.25 31ZM20 7a4 4 0 1 1-8 0V3H9v8a3 3 0 0 1-3 3v12.672L23.014 7.735L23 3h-3v4ZM8.143 29H26V14a2.988 2.988 0 0 1-2.756-1.808L8.143 29Z"/>`),
		g.Group(children),
	)
}