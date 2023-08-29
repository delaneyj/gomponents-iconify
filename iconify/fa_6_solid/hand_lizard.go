package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandLizard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 112c0-26.5 21.5-48 48-48h238.5c36.8 0 71.2 18 92.1 48.2l113.5 164c13 18.7 19.9 41 19.9 63.8v76c0 17.7-14.3 32-32 32h-96c-17.7 0-32-14.3-32-32v-13.8L273.9 352H112c-26.5 0-48-21.5-48-48s21.5-48 48-48h128c26.5 0 48-21.5 48-48s-21.5-48-48-48H48c-26.5 0-48-21.5-48-48z"/>`),
		g.Group(children),
	)
}