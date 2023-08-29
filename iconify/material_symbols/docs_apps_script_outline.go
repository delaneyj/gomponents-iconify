package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocsAppsScriptOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21q-.425 0-.713-.288T12 20q0-.425.288-.713T13 19h2.825l-2.3-1.625q-.35-.25-.413-.638t.163-.737q.225-.35.625-.413t.75.163l2.325 1.6L16 14.7q-.15-.375.025-.75t.575-.525q.4-.15.775.025t.525.575l.95 2.65l.75-2.725q.125-.4.462-.612t.738-.088q.4.125.625.463t.1.737l-1.55 5.8q-.1.35-.363.55T19 21h-6ZM3 18v-2h7.075q-.075.525-.063 1t.088 1H3Zm0-4v-2h9.65q-.575.4-1.038.9T10.8 14H3Zm0-4V8h15v2H3Zm0-4V4h15v2H3Z"/>`),
		g.Group(children),
	)
}