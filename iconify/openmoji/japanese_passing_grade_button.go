package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapanesePassingGradeButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiJapanesePassingGradeButton0" stroke-linecap="round" stroke-linejoin="round" d="M20 31c7 0 15-11 15-11s9.5 10.5 15.5 11.5m-22-1h14m-17 20v-13h20v13m-19-2h19"/></defs><path fill="#ea5a47" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-miterlimit="10" d="M59.035 60h-46.07a.968.968 0 0 1-.965-.965v-46.07a.968.968 0 0 1 .965-.965h46.07a.968.968 0 0 1 .965.965v46.07a.968.968 0 0 1-.965.965Z"/><use href="#openmojiJapanesePassingGradeButton0" stroke-linecap="round" stroke-linejoin="round"/></g><use href="#openmojiJapanesePassingGradeButton0" fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>`),
		g.Group(children),
	)
}