package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sofa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSofa0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M12 21H4v14h8V21Zm32 0h-8v14h8V21Z"/><path stroke-linecap="round" d="M36 27H12v8h24v-8ZM8 20V8h32v12M8 36v4m32-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSofa0)"/>`),
		g.Group(children),
	)
}