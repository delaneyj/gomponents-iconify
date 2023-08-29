package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 576L832 683v85L512 960L192 768v-85L0 576l192-192L0 192L320 0l192 192L704 0l320 192l-192 192zM512 243L235 410l277 166l277-166z"/>`),
		g.Group(children),
	)
}