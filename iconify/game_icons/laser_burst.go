package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaserBurst(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m291.17 299.252l109.093 150.173l-21.76 15.786l-109.07-150.16v72.135h-26.865V315.05l-109.07 150.16l-21.74-15.795l109.07-150.163l-68.6 22.287l-8.307-25.555l68.602-22.287L36 216.348l8.307-25.565l176.533 57.362l-42.404-58.36l21.686-15.776l42.446 58.36V46.79h26.864v185.58l42.446-58.36l21.74 15.795l-42.404 58.36l176.48-57.382L476 216.337L299.467 273.71l68.602 22.286l-8.308 25.554z"/>`),
		g.Group(children),
	)
}