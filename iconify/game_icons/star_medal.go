package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarMedal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M41 27.7v46h430v-46zm32 64v112.5l62 20.7V91.7zm80 0v139.2l79.3 26.4c-10.4 7.5-17.3 19.7-17.3 33.4c0 21 16.1 38.5 36.5 40.7l-22.3 43.7l-58.2 9.3l41.6 41.7l-9.1 58.2l52.5-26.7l52.5 26.7l-9.1-58.2l41.6-41.7l-58.2-9.3l-22.3-43.7c20.4-2.2 36.5-19.7 36.5-40.7c0-13.7-6.9-25.9-17.3-33.4l79.3-26.4V91.7h-94v159c-2.9-.6-5.9-1-9-1c-3.1 0-6.1.4-9 1v-159zm224 0v133.2l62-20.7V91.7zm-121 176c12.8 0 23 10.2 23 23s-10.2 23-23 23s-23-10.2-23-23s10.2-23 23-23z"/>`),
		g.Group(children),
	)
}