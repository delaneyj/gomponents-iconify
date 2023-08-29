package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retweet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 384L832 640L640 384h128V128H384L288 0h608v384h128zM256 256v256h384l96 128H128V256H0L192 0l192 256H256z"/>`),
		g.Group(children),
	)
}