package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PollTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.752 2a2.752 2.752 0 0 1 2.752 2.752V19.25a2.752 2.752 0 1 1-5.504 0V4.752A2.752 2.752 0 0 1 11.752 2Zm7 5a2.752 2.752 0 0 1 2.752 2.752v9.498a2.752 2.752 0 1 1-5.504 0V9.752A2.752 2.752 0 0 1 18.752 7Zm-14 5a2.752 2.752 0 0 1 2.752 2.752v4.498a2.752 2.752 0 0 1-5.504 0v-4.498A2.752 2.752 0 0 1 4.752 12Z"/>`),
		g.Group(children),
	)
}