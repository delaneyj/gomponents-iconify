package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BirthdayCake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBirthdayCake0"><g fill="none"><path fill="#555" d="M8 40h32V24H8v16Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 40H8m32 0H4h4m32 0h4m-4 0V24H8v16"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m40 34l-4-2l-4 2l-4-2l-4 2l-4-2l-4 2l-4-2l-4 2m24-10v-9m-8 9v-9m-8 9v-9m16-5V8m-8 2V8m-8 2V8M8 24v16m32-16v16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBirthdayCake0)"/>`),
		g.Group(children),
	)
}