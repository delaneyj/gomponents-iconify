package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.313 4V2.144C16.313.96 15.353 0 14.169 0H7.831A2.142 2.142 0 0 0 5.69 2.189v-.002V4H0v2h.575c.196.023.372.099.515.214l-.002-.002c.119.157.203.346.239.552l.001.008l1.187 15.106c.094 1.84.094 2.118 2.25 2.118h12.462c2.16 0 2.16-.275 2.25-2.113l1.187-15.1c.036-.217.12-.409.242-.572l-.002.003a.994.994 0 0 1 .508-.212h.58v-2h-5.687zM7 2.187c0-.6.487-.938 1.106-.938h5.734c.618 0 1.162.344 1.162.938V4h-8zM6.469 20l-.64-12h1.269l.656 12zm5.225 0H10.32V8h1.375zm3.85 0h-1.275l.656-12h1.269z"/>`),
		g.Group(children),
	)
}