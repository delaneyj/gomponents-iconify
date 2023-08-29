package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chopsticks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.9 8.93L8.715 29.354a.87.87 0 0 1-1.69-.42l4.6-20.654l2.277.65Zm.247-.97l1.107-4.356c.25-.94-.61-1.8-1.55-1.55c-.45.12-.79.48-.9.93l-.962 4.319l2.305.657Zm7.443 3.35L11.375 29.564a.89.89 0 0 1-.76.44c-.66 0-1.08-.7-.77-1.29l9.62-18.435l2.127 1.03Zm-1.663-1.918l2.217-4.248c.22-.41.65-.67 1.11-.67c.97 0 1.58 1.05 1.1 1.9l-2.274 4.061l-2.153-1.043Z"/>`),
		g.Group(children),
	)
}