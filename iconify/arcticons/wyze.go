package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wyze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-5.456 22.956H39.5m-4.456-8.912H39.5M35.044 24h2.905m-2.905-4.456v8.912"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.413 19.544l-2.229 8.912l-2.228-8.912l-2.228 8.912L8.5 19.544m16.763 0L22.31 24l-2.952-4.456m2.952 8.912V24m4.893-4.456h5.904l-5.904 8.912h5.904"/>`),
		g.Group(children),
	)
}