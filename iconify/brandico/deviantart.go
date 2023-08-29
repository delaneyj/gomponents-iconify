package brandico

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="m362.844 293.375l4.906 339.219H161.281c23.034-97.986 69.423-176.972 203-196.656l-1.438-44.25C69.007 430.08 6.665 591.516 2.53 706.626l473.375-.281V423.189c142.178-3.648 319.852 1.934 368.688 207.719l-210.313-1.719V484.908c-37.66-10.995-70.11-11.183-106.031-8.687l-1.75 229.531l469.313-.094c4.511-229.461-199.386-333.433-519.906-328.5v-83.781H362.843z"/>`),
		g.Group(children),
	)
}