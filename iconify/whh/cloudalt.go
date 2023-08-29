package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 768H224q-93 0-158.5-65.5T0 544q0-84 55-147t137-74v-3q0-87 43-160.5T351.5 43T512 0q100 0 180.5 56T809 202q94 24 154.5 101.5T1024 480q0 111-74 193.5T768 768z"/>`),
		g.Group(children),
	)
}