package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmWatsonOpenscale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 28c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3zm0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zm24 4c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3zm0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zM22.4 4.4l-.9 1.8C26.1 8.4 29 13 29 18c0 .7-.1 1.4-.2 2.1l2 .3c.1-.8.2-1.6.2-2.5c0-5.7-3.4-11-8.6-13.5zM16 7c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3zm0-4c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1z"/><path fill="currentColor" d="m25.5 13.6l-1-1.7l-7.5 4.4V9h-2v7.2L6.2 11c-.5-.3-1.1-.1-1.4.3s-.1 1.1.3 1.4L14 18l-6.2 3.6l1 1.7l6.2-3.6V30c0 .6.4 1 1 1s1-.4 1-1V19.8l6.2 3.7l1-1.7L18 18l7.5-4.4z"/>`),
		g.Group(children),
	)
}