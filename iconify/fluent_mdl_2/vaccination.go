package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vaccination(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1955 637l-99-99l-166 166l99 99l-90 90l-99-99l-869 870H474l-202 203l-272 90l384-384v-256l870-869l-99-99l90-90l99 99l166-166l-99-99l90-90l544 544l-90 90zM677 1536l832-832l-165-165l-229 229l82 83l-90 90l-83-82l-102 101l83 83l-90 90l-83-83l-101 102l82 83l-90 90l-83-82l-128 128v165h165zm923-923l165-165l-165-165l-165 165l165 165z"/>`),
		g.Group(children),
	)
}