package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M224 184h128m-128 72h128m-128 71h128"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32" d="M448 258c0-106-86-192-192-192S64 152 64 258s86 192 192 192s192-86 192-192Z"/><circle cx="168" cy="184" r="8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"/><circle cx="168" cy="257" r="8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"/><circle cx="168" cy="328" r="8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"/>`),
		g.Group(children),
	)
}