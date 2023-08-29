package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagChile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#1F429B" d="M13 5H4a4 4 0 0 0-4 4v9h13V5zm-4.663 9.292l-1.882-1.367l-1.882 1.367l.719-2.212l-1.882-1.368h2.326L6.455 8.5l.719 2.212H9.5L7.618 12.08l.719 2.212z"/><path fill="#EEE" d="M32 5H13v13h23V9a4 4 0 0 0-4-4z"/><path fill="#D42D27" d="M0 18v9a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-9H0z"/><path fill="#FFF" d="M7.174 10.712L6.455 8.5l-.719 2.212H3.41l1.882 1.368l-.719 2.212l1.882-1.367l1.882 1.367l-.719-2.212L9.5 10.712z"/>`),
		g.Group(children),
	)
}