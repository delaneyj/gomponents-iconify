package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActivityFeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1088V768H0V128h1664v640H448l-320 320zm0-448h128v139l139-139h1141V256H128v384zm1920 384v640h-128v320l-320-320H384v-640h1664zm-128 128H512v384h1141l139 139v-139h128v-384z"/>`),
		g.Group(children),
	)
}