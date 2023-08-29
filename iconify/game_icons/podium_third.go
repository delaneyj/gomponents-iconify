package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodiumThird(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M417.945 168.902c-11.593 0-23 12.8-23 31c0 18.201 11.407 31 23 31c11.594 0 23-12.799 23-31c0-18.2-11.406-31-23-31zm-103.95 2.975l-16.099 8.05c15.093 30.185 51.37 56.81 82.188 74.442l14.195 184.533h14.666v-103h18v103h14.666l11.854-154.093l13.928 51.892l17.382-4.664c-6.156-34.54-15.319-97.864-34.212-102.39c-7.307 11.535-18.869 19.255-32.618 19.255c-13.749 0-25.31-7.72-32.617-19.256c-26.443-8.62-51.424-27.87-71.334-57.769zM169 313v96H25v78h462v-30H343V313H169z"/>`),
		g.Group(children),
	)
}