package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flatbread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="m6.69 17.07l.18.18c.24.25.63.25.88 0l9.51-9.51c.24-.24.24-.64 0-.88l-.18-.18a.628.628 0 0 0-.88 0l-9.51 9.51c-.24.24-.24.64 0 .88Zm1.41 5.77c-.29.29-.29.77 0 1.06c.29.29.77.29 1.06 0L23.9 9.16c.29-.29.29-.77 0-1.06a.754.754 0 0 0-1.06 0L8.1 22.84Zm6.83 2.48l-.18-.18a.628.628 0 0 1 0-.88l9.51-9.51c.24-.24.64-.24.88 0l.18.18c.24.24.24.64 0 .88l-9.51 9.51c-.25.24-.64.24-.88 0Z"/><path d="M15 1C7.268 1 1 7.268 1 15v2c0 7.732 6.268 14 14 14h2c7.732 0 14-6.268 14-14v-2c0-7.732-6.268-14-14-14h-2ZM3 15C3 8.373 8.373 3 15 3h2c6.627 0 12 5.373 12 12v2c0 6.627-5.373 12-12 12h-2C8.373 29 3 23.627 3 17v-2Z"/></g>`),
		g.Group(children),
	)
}