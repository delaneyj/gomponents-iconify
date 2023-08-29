package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocsAppsScriptOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21q-.425 0-.713-.288T12 20q0-.425.288-.713T13 19h2.825l-2.3-1.625q-.35-.25-.413-.638t.163-.737q.225-.35.625-.413t.75.163l2.325 1.6L16 14.7q-.15-.375.025-.75t.575-.525q.4-.15.775.025t.525.575l.95 2.65l.75-2.725q.125-.4.462-.612t.738-.088q.4.125.625.463t.1.737l-1.55 5.8q-.1.35-.363.55T19 21h-6Zm-9-3q-.425 0-.713-.288T3 17q0-.425.288-.713T4 16h6.075q-.075.525-.063 1t.088 1H4Zm0-4q-.425 0-.713-.288T3 13q0-.425.288-.713T4 12h8.65q-.575.4-1.038.9T10.8 14H4Zm0-4q-.425 0-.713-.288T3 9q0-.425.288-.713T4 8h13q.425 0 .713.288T18 9q0 .425-.288.713T17 10H4Zm0-4q-.425 0-.713-.288T3 5q0-.425.288-.713T4 4h13q.425 0 .713.288T18 5q0 .425-.288.713T17 6H4Z"/>`),
		g.Group(children),
	)
}