package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicalNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00b8f0" d="M12.93 6.86a1 1 0 0 0-.93 1v21.8a5.94 5.94 0 0 0-4-.72a6 6 0 0 0 1.26 11.9A6.15 6.15 0 0 0 15 34.6V15.71l23-1.66v13.17a5.94 5.94 0 0 0-4-.72a6 6 0 0 0 1.26 11.91A6.17 6.17 0 0 0 41 32.17V5.91a1 1 0 0 0-1.07-1Z"/><path fill="#45413c" d="M12 45.5a12 1.5 0 1 0 24 0a12 1.5 0 1 0-24 0Z" opacity=".15"/><path fill="#4acfff" d="M34 29.54a6 6 0 0 1 3.27.35a.54.54 0 0 0 .75-.5v-2.17a5.94 5.94 0 0 0-4-.72a6 6 0 0 0-4.92 5a6.13 6.13 0 0 0 .13 2.39A6 6 0 0 1 34 29.54ZM8 32a6.06 6.06 0 0 1 3.27.36a.52.52 0 0 0 .51-.06a.54.54 0 0 0 .24-.45v-2.2a5.94 5.94 0 0 0-4-.72a6 6 0 0 0-4.79 7.42A6.06 6.06 0 0 1 8 32ZM39.93 4.92l-27 1.94a1 1 0 0 0-.93 1v3a1 1 0 0 1 .93-1l27-1.86A1 1 0 0 1 41 9V6a1 1 0 0 0-1.07-1.08Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M12.93 6.86a1 1 0 0 0-.93 1v21.8a5.94 5.94 0 0 0-4-.72a6 6 0 0 0 1.26 11.9A6.15 6.15 0 0 0 15 34.6V15.71l23-1.66v13.17a5.94 5.94 0 0 0-4-.72a6 6 0 0 0 1.26 11.91A6.17 6.17 0 0 0 41 32.17V5.91a1 1 0 0 0-1.07-1Z"/>`),
		g.Group(children),
	)
}