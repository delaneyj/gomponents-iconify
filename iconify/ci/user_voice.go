package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserVoice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19c0-2.21-2.686-4-6-4s-6 1.79-6 4M16.828 5.172a3.999 3.999 0 0 1 0 5.657M19 3a7.07 7.07 0 0 1 0 10M9 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8Z"/>`),
		g.Group(children),
	)
}