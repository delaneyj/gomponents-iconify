package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWindows0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="3.833" d="m6.75 11.063l12.938-1.725v12.075H6.75v-10.35Zm18.112-2.218L41.25 6.75v14.663H24.862V8.845Zm0 18.605l16.388.383V41.25l-16.388-2.683V27.45ZM6.75 26.587l12.938.312V37.8L6.75 35.62v-9.032Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWindows0)"/>`),
		g.Group(children),
	)
}