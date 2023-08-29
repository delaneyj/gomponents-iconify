package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRulerOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M43 41H23"/><path fill="#555" stroke-linejoin="round" d="M38.718 5H21L5 41h17.662L38.718 5Z"/><path stroke-linejoin="round" d="M9.959 29.882h8.028m-4.722-7.412h8.028m-4.519-7.87h8.029"/><path d="M21 5L5 41"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRulerOne0)"/>`),
		g.Group(children),
	)
}