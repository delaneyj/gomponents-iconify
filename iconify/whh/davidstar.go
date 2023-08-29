package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Davidstar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M897 768H608l-160 256l-160-256H0l128-256L0 256h288L448 0l160 256h289L768 512zm-97-64l-67-135l-85 135h152zM448 896l86-128H363zM96 704h152l-84-135zm0-384l68 135l84-135H96zm352-192l-85 128h171zm128 192H320L192 512l128 192h256l128-192zm224 0H648l85 135z"/>`),
		g.Group(children),
	)
}