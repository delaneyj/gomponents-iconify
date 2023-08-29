package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M384 368h24a40.12 40.12 0 0 0 40-40V168a40.12 40.12 0 0 0-40-40H104a40.12 40.12 0 0 0-40 40v160a40.12 40.12 0 0 0 40 40h24"/><rect width="256" height="208" x="128" y="240" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="24.32" ry="24.32"/><path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M384 128v-24a40.12 40.12 0 0 0-40-40H168a40.12 40.12 0 0 0-40 40v24"/><circle cx="392" cy="184" r="24" fill="currentColor"/>`),
		g.Group(children),
	)
}