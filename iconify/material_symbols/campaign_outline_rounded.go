package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CampaignOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 13q-.425 0-.713-.288T18 12q0-.425.288-.713T19 11h2q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-2Zm-.6 6.4l-1.6-1.2q-.35-.25-.4-.65t.2-.75q.25-.35.65-.4t.75.2l1.6 1.2q.35.25.4.65t-.2.75q-.25.35-.65.4t-.75-.2Zm-.4-12q-.35.25-.75.2t-.65-.4q-.25-.35-.2-.75t.4-.65l1.6-1.2q.35-.25.75-.2t.65.4q.25.35.2.75t-.4.65L18 7.4ZM5 19v-4H4q-.825 0-1.413-.588T2 13v-2q0-.825.588-1.413T4 9h4l3.475-2.1q.5-.3 1.012 0t.513.875v8.45q0 .575-.513.875t-1.012 0L8 15H7v4H5Zm2.5-7Zm6.5 3.35v-6.7q.675.6 1.088 1.463T15.5 12q0 1.025-.413 1.888T14 15.35Zm-3-.9v-4.9L8.55 11H4v2h4.55L11 14.45Z"/>`),
		g.Group(children),
	)
}