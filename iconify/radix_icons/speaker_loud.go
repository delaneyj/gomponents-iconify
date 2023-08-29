package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerLoud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.47 1.05a.5.5 0 0 1 .28.45v12a.5.5 0 0 1-.807.395L3.221 11H1.5A1.5 1.5 0 0 1 0 9.5v-4A1.5 1.5 0 0 1 1.5 4h1.721l3.722-2.895a.5.5 0 0 1 .527-.054Zm-.72 1.472L3.7 4.895A.5.5 0 0 1 3.393 5H1.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h1.893a.5.5 0 0 1 .307.105l3.05 2.373V2.522Zm3.528 1.326a.4.4 0 0 1 .555.111a6.407 6.407 0 0 1 0 7.081a.4.4 0 0 1-.666-.443a5.607 5.607 0 0 0 0-6.194a.4.4 0 0 1 .111-.555Zm2.4-2.418a.4.4 0 0 0-.61.518a8.602 8.602 0 0 1 0 11.104a.4.4 0 0 0 .61.518a9.402 9.402 0 0 0 0-12.14Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}