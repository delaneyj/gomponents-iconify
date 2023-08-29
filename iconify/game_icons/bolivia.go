package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolivia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M43.896 68.829c29.684-4.985 63.931-12.904 148.312-53.051l32.89 102.214L369.21 161.28l7.59 96.475l61.805 3.512l29.498 72.34l-25.284 58.996c-32.53-13.45-59.839-33.256-121.503-11.238c-16.164 19.127-23.681 57.278-33.712 89.899c-76.82-14.398-136.704-9.87-168.56 25.986L50.92 293.575c10.326-68.28 10.417-141.922-7.023-224.746z"/>`),
		g.Group(children),
	)
}