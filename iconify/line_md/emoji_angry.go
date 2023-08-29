package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiAngry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="60" stroke-dashoffset="60" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M8 16C8.5 15 9.79086 14 12 14C14.2091 14 15.5 15 16 16"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;0"/></path><g stroke-dasharray="5" stroke-dashoffset="5" stroke-width="1"><path d="M9.5 8.5L7.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="5;0"/></path><path d="M14.5 8.5L16.5 7.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.8s" dur="0.2s" values="5;0"/></path></g></g><g fill="currentColor" fill-opacity="0"><ellipse cx="9" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		g.Group(children),
	)
}