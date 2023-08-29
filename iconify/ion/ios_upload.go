package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M264 144v162h-16V144H96v304h320V144z" fill="currentColor"/><path d="M264 63.4l54.8 54.7 11.6-11.6L256 32l-74.5 74.5 11.7 11.6L248 63.4V144h16z" fill="currentColor"/>`),
		g.Group(children),
	)
}