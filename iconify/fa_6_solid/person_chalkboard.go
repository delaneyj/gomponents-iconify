package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonChalkboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 96a48 48 0 1 0 0-96a48 48 0 1 0 0 96zm-8 384V352h16v128c0 17.7 14.3 32 32 32s32-14.3 32-32V192h136c17.7 0 32-14.3 32-32s-14.3-32-32-32h-16V64h192v192H384v-32h-64v48c0 26.5 21.5 48 48 48h224c26.5 0 48-21.5 48-48V48c0-26.5-21.5-48-48-48H368c-26.5 0-48 21.5-48 48v80H177.1c-33.7 0-64.9 17.7-82.3 46.6l-58.3 97c-9.1 15.1-4.2 34.8 10.9 43.9s34.8 4.2 43.9-10.9l28.7-47.7V480c0 17.7 14.3 32 32 32s32-14.3 32-32z"/>`),
		g.Group(children),
	)
}