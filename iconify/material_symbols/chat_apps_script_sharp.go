package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatAppsScriptSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.45 20q-.425 0-.712-.288T13.45 19q0-.425.288-.713T14.45 18h2.825l-2.3-1.625q-.35-.25-.412-.638t.162-.737q.225-.35.625-.413t.75.163l2.325 1.6l-.975-2.65q-.15-.375.025-.75t.575-.525q.4-.15.775.025t.525.575l.95 2.65l.75-2.725q.125-.4.462-.612t.738-.088q.4.125.613.463t.087.737L21.2 20h-6.75ZM7 9h8V7H7v2Zm0 4h5v-2H7v2Zm-4 7V3h16v7.075q-.25-.05-.5-.063T18 10q-2.525 0-4.263 1.75T12 16q0 .25.013.5t.062.5H6l-3 3Z"/>`),
		g.Group(children),
	)
}