package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagHonduras(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#EEE" d="M0 14h36v8H0z"/><path fill="#0156A3" d="M32 5H4a4 4 0 0 0-4 4v5h36V9a4 4 0 0 0-4-4zM0 27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-5H0v5z"/><path fill="#2657A7" d="m18.265 17.76l-.249-.766l-.249.766h-.806l.652.473l-.249.767l.652-.474l.652.474l-.249-.767l.652-.473zm5.332 2.48l-.249.766l.652-.473l.652.473l-.249-.766l.652-.473h-.806L24 19l-.249.767h-.806zm.652-4.48L24 14.994l-.249.766h-.806l.652.473l-.249.767l.652-.474l.652.474l-.249-.767l.652-.473zm-12.707 4.48l-.249.766l.652-.473l.652.473l-.249-.766l.652-.473h-.806L11.945 19l-.249.767h-.806zm.652-4.48l-.249-.766l-.249.766h-.806l.652.473l-.249.767l.652-.474l.652.474l-.249-.767l.652-.473z"/>`),
		g.Group(children),
	)
}