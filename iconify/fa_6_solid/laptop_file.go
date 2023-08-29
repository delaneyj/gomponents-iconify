package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 0C92.7 0 64 28.7 64 64v224H19.2C8.6 288 0 296.6 0 307.2C0 349.6 34.4 384 76.8 384H320v-96H128V64h320v32h64V64c0-35.3-28.7-64-64-64H128zm384 128H400c-26.5 0-48 21.5-48 48v288c0 26.5 21.5 48 48 48h192c26.5 0 48-21.5 48-48V256h-96c-17.7 0-32-14.3-32-32v-96zm32 0v96h96l-96-96z"/>`),
		g.Group(children),
	)
}