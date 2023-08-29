package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Igloo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 33.8V160H48.5C100.2 82.8 188.1 32 288 32c10.8 0 21.5.6 32 1.8zM352 160V39.1C424.9 55.7 487.2 99.8 527.5 160H352zM29.9 192H96v128H0c0-46 10.8-89.4 29.9-128zM192 320h-64V192h320v128h-64v32h192v80c0 26.5-21.5 48-48 48H352V352c0-35.3-28.7-64-64-64s-64 28.7-64 64v128H48c-26.5 0-48-21.5-48-48v-80h192v-32zm288 0V192h66.1c19.2 38.6 29.9 82 29.9 128h-96z"/>`),
		g.Group(children),
	)
}