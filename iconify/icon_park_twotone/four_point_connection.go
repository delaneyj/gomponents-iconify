package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourPointConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFourPointConnection0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 28v12h12m8 0h12V28m0-8V8H28m-8 0H8v12"/><path fill="#555" d="M44 20h-8v8h8v-8Zm-32 0H4v8h8v-8Zm16 16h-8v8h8v-8Zm0-32h-8v8h8V4Z"/><path d="M24 18v12m-6-6h12M28 8l12 12M20 8L8 20m12 20L8 28m32 0L29 40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFourPointConnection0)"/>`),
		g.Group(children),
	)
}