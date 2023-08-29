package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorWatcher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M105 25v462h302V25H105zm44 28h214v150H149V53zm18 18v114h178V71H167zm16 16h146v82H183V87zm15.8 25c.8 7.9 5.2 18 11.1 24.8c11.2-.2 25.9-3.4 36.1-8.8c-21-2.3-38.9-9.3-47.2-16zm114.4 0c-8.3 6.7-26.2 13.7-47.2 16c10.2 5.4 24.9 8.6 36.1 8.8c5.9-6.8 10.3-16.9 11.1-24.8zM256 279h128v18h-17v14h17v18h-64v-18h29v-14h-93v-18z"/>`),
		g.Group(children),
	)
}