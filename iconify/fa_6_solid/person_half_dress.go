package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonHalfDress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M160 0a48 48 0 1 1 0 96a48 48 0 1 1 0-96zm8 352V128h6.9c33.7 0 64.9 17.7 82.3 46.6l58.3 97c9.1 15.1 4.2 34.8-10.9 43.9s-34.8 4.2-43.9-10.9L232 256.9V480c0 17.7-14.3 32-32 32s-32-14.3-32-32V352zM58.2 182.3c19.9-33.1 55.3-53.5 93.8-54.3v352c0 17.7-14.3 32-32 32s-32-14.3-32-32v-96H70.2c-10.9 0-18.6-10.7-15.2-21.1l38.3-114.8l-33.9 56.4c-9.1 15.1-28.8 20-43.9 10.9s-20-28.8-10.9-43.9l53.6-89.2z"/>`),
		g.Group(children),
	)
}