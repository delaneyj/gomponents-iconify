package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilDrum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.02 2c.54 0 .98.44.98.98V4h1c.55 0 1 .45 1 1s-.45 1-1 1v6c.55 0 1 .45 1 1s-.45 1-1 1v6c.55 0 1 .45 1 1s-.45 1-1 1v6c.55 0 1 .45 1 1s-.45 1-1 1H7c-.55 0-1-.45-1-1s.45-1 1-1v-6c-.55 0-1-.45-1-1s.45-1 1-1v-6c-.55 0-1-.45-1-1s.45-1 1-1V6c-.55 0-1-.45-1-1s.45-1 1-1h12V2.98c0-.54.44-.98.99-.98h3.03ZM24 6h-2v6h2V6Zm0 8h-2v6h2v-6Zm0 8h-2v6h2v-6Zm-3 6v-6h-1v6h1Zm-1-8h1v-6h-1v6Zm0-8h1V6h-1v6Z"/>`),
		g.Group(children),
	)
}