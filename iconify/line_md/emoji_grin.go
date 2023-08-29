package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiGrin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdEmojiGrin0"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-width="2" d="M8 14C8.5 15.5 9.79086 17 12 17C14.2091 17 15.5 15.5 16 14M8 14C9 14.5 10 15 12 15C14 15 15 14.5 16 14"/></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/></path><g fill="currentColor"><rect width="0" height="7" x="6" y="12" mask="url(#lineMdEmojiGrin0)"><animate fill="freeze" attributeName="width" begin="1s" dur="0.2s" values="0;12"/></rect><ellipse cx="9" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.2s" values="0;1"/></ellipse><ellipse cx="15" cy="9.5" fill-opacity="0" rx="1" ry="1.5"><animate fill="freeze" attributeName="fill-opacity" begin="0.8s" dur="0.2s" values="0;1"/></ellipse></g>`),
		g.Group(children),
	)
}