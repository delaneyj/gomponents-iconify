package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percentage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm5 2c-1.645 0-3 1.355-3 3v1c0 1.645 1.355 3 3 3s3-1.355 3-3v-1c0-1.645-1.355-3-3-3zm7.594 0L10 23h2.406L22 9zM12 11c.566 0 1 .434 1 1v1c0 .566-.434 1-1 1c-.566 0-1-.434-1-1v-1c0-.566.434-1 1-1zm8 5c-1.645 0-3 1.355-3 3v1c0 1.645 1.355 3 3 3s3-1.355 3-3v-1c0-1.645-1.355-3-3-3zm0 2c.566 0 1 .434 1 1v1c0 .566-.434 1-1 1c-.566 0-1-.434-1-1v-1c0-.566.434-1 1-1z"/>`),
		g.Group(children),
	)
}