package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 8h7.1a5.002 5.002 0 0 1 9.8 0H28a3 3 0 0 1 3 3v17a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V11a3 3 0 0 1 3-3Zm14.83 0a3.001 3.001 0 0 0-5.66 0h5.66ZM3 15.44l.032.001A3 3 0 0 0 6 18h6a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2h6a3 3 0 0 0 2.968-2.559H29V11a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v4.44ZM19 19v-1a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v1h6Zm-6 1v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-1h-6Zm-1 0H6a4.977 4.977 0 0 1-3-1v9a1 1 0 0 0 1 1h24a1 1 0 0 0 1-1v-9c-.835.628-1.874 1-3 1h-6v1a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-1Z"/>`),
		g.Group(children),
	)
}