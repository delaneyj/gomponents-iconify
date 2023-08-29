package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSClear0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" stroke-linecap="round" d="M20 5.914h8v8h15v8H5v-8h15v-8Z" clip-rule="evenodd"/><path fill="#fff" stroke="#fff" d="M8 40h32V22H8v18Z"/><path stroke="#000" stroke-linecap="round" d="M16 39.898v-5.984m8 5.984v-6m8 6v-5.984"/><path stroke="#fff" stroke-linecap="round" d="M12 40h24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSClear0)"/>`),
		g.Group(children),
	)
}