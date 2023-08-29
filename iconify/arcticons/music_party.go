package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicParty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.116 11.774v24.452M8.804 16.833v14.334M4.5 21.471v5.058m12.957-9.696v14.334m4.49-9.696v5.058m12.722-14.755v24.452m-4.312-19.393v14.334m-4.304-9.696v5.058m12.957-9.696v14.334m4.49-9.696v5.058"/>`),
		g.Group(children),
	)
}