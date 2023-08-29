package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosUploadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M288 144v16h112v272H112V160h112v-16H96v304h320V144z" fill="currentColor"/><path d="M193.1 118.1l-11.6-11.6L256 32l74.5 74.5-11.6 11.6-54.7-54.7v243h-16.4v-243z" fill="currentColor"/>`),
		g.Group(children),
	)
}