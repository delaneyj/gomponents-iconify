package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M151 16c-14.774 0-30 15.226-30 30v105c0 14.774 11.946 26.718 26.718 26.718H181V334.28h-33.282c-14.773 0-26.718 11.946-26.718 26.718v105c0 14.774 15.227 30 30 30h30v-90l75-45l75 45v90h30c14.774 0 30-15.226 30-30v-105c0-14.773-11.946-26.718-26.718-26.718H331V177.718h33.282C379.056 177.718 391 165.772 391 151V46c0-14.773-15.226-30-30-30h-30v90l-75 45l-75-45V16h-30z"/>`),
		g.Group(children),
	)
}