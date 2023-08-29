package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ela(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#3FBADF"/><g fill="#FFF" fill-rule="nonzero"><path fill-opacity=".4" d="m11 22.119l5-2.82v5.635l-5-2.815zm0-9l5-2.82v5.635l-5-2.815z"/><path fill-opacity=".7" d="m26 19.23l-5 2.886V16.43l5 2.8zm0-9l-5 2.886V7.43l5 2.8z"/><path fill-opacity=".8" d="M11 22.116v-5.683l5 2.87l-5 2.813zm0-9V7.433l5 2.87l-5 2.813z"/><path d="m21 22.116l-5-2.812l5-2.874v5.686zm0-9l-5-2.812l5-2.874v5.686z"/><path fill-opacity=".6" d="m21 22.116l-5 2.818v-5.63l5 2.812zm0-9l-5 2.818v-5.63l5 2.812z"/><path fill-opacity=".5" d="M11 16.433v5.683l-5-2.885l5-2.798zm0-9v5.683l-5-2.885l5-2.798z"/></g></g>`),
		g.Group(children),
	)
}