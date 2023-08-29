package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparkle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M35.106 18.815s-.754 16.324-16.128 16.37a.648.648 0 1 0 0 1.296h.004s16.324 1.1 16.37 16.474v.147a.648.648 0 0 0 1.296 0c.046-15.375 16.37-16.62 16.37-16.62h.005a.648.648 0 1 0 0-1.297c-15.375-.045-16.62-16.37-16.62-16.37a.648.648 0 0 0-1.297 0Z"/><path fill="#5c9e31" d="M60 61H12a.945.945 0 0 1-1-1V12a.945.945 0 0 1 1-1h48a.945.945 0 0 1 1 1v48a.945.945 0 0 1-1 1Z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M35.106 18.857s-.754 16.324-16.128 16.37a.648.648 0 1 0 0 1.296h.004s16.324 1.1 16.37 16.474v.146a.648.648 0 0 0 1.296 0c.046-15.375 16.37-16.62 16.37-16.62h.005a.648.648 0 1 0 0-1.296c-15.375-.046-16.62-16.37-16.62-16.37a.648.648 0 0 0-1.297 0Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M60 60.958H12a.945.945 0 0 1-1-1v-48a.945.945 0 0 1 1-1h48a.945.945 0 0 1 1 1v48a.945.945 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}