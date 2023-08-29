package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundMute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m2.47 13.513l14-11.142l.873 1.097l-14 11.142zm2.483-3.601V5.01H3.128c-1.283 0-2.115 1.084-2.115 2.46v1.024c0 1.459.769 2.459 2.115 2.459h.386l1.439-1.041zm2.924 2.747l5.059 3.313c.586 0 1.06-.4 1.06-.895V7.919l-6.119 4.74zM13.987.971c0-.507-.499-.92-1.115-.92L7.114 3.73C6.499 3.73 6 4.142 6 4.65v4.189l7.987-6.243V.971z"/>`),
		g.Group(children),
	)
}