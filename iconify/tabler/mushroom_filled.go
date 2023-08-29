package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MushroomFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M15 15v4a3 3 0 0 1-5.995.176L9 19v-4h6zM4.9 13a1.9 1.9 0 0 1-1.894-1.752L3 11.1C3 6.077 7.027 2 12 2s9 4.077 9 9.1a1.9 1.9 0 0 1-1.752 1.894L19.1 13H4.9z"/></g>`),
		g.Group(children),
	)
}