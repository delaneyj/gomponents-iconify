package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdUploadOffOutline0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M6 19h12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.4s" values="14;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M12 15 h2 v-6 h2.5 L12 4.5M12 15 h-2 v-6 h-2.5 L12 4.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="18;0"/></path><g stroke-dasharray="26" stroke-dashoffset="26" transform="rotate(45 13 12)"><path stroke="#000" d="M0 11h24"/><path d="M1 13h22"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="26;0"/></g></g></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdUploadOffOutline0)"/>`),
		g.Group(children),
	)
}