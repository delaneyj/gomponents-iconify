package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkWired(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 64h128v64H256V64zM240 0c-26.5 0-48 21.5-48 48v96c0 26.5 21.5 48 48 48h48v32H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h96v32H80c-26.5 0-48 21.5-48 48v96c0 26.5 21.5 48 48 48h160c26.5 0 48-21.5 48-48v-96c0-26.5-21.5-48-48-48h-48v-32h256v32h-48c-26.5 0-48 21.5-48 48v96c0 26.5 21.5 48 48 48h160c26.5 0 48-21.5 48-48v-96c0-26.5-21.5-48-48-48h-48v-32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32H352v-32h48c26.5 0 48-21.5 48-48V48c0-26.5-21.5-48-48-48H240zM96 448v-64h128v64H96zm320-64h128v64H416v-64z"/>`),
		g.Group(children),
	)
}