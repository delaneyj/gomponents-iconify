package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M25 31.041a5.69 5.69 0 0 0-8 0L21 35l4-3.959zm8-7.921c-6.63-6.562-17.37-6.562-24 0l4 3.959c4.42-4.37 11.58-4.37 16 0l4-3.959zm8-7.921c-11.05-10.932-28.95-10.932-40 0l4 3.959c8.84-8.751 23.16-8.751 32 0l4-3.959z"/>`),
		g.Group(children),
	)
}