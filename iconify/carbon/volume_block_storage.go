package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeBlockStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 21h-7v-7h7zm-5-2h3v-3h-3zm5 11h-7v-7h7zm-5-2h3v-3h-3zm14-7h-7v-7h7zm-5-2h3v-3h-3zm5 11h-7v-7h7zm-5-2h3v-3h-3z"/><path fill="currentColor" d="M8 28H4a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h7.586A1.986 1.986 0 0 1 13 4.586L16.414 8H28a2.002 2.002 0 0 1 2 2v8h-2v-8H15.586l-4-4H4v20h4Z"/>`),
		g.Group(children),
	)
}