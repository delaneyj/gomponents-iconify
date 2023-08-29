package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicalNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.41 2.869c0-.53-.46-.93-.98-.86l-16.97 2.2c-.43.06-.76.43-.76.87v13.61c-.59-.2-1.23-.32-1.89-.32c-3.21 0-5.81 2.61-5.81 5.82c0 3.21 2.6 5.81 5.81 5.81c3.21 0 5.81-2.6 5.81-5.81c0-.12-.01-.25-.01-.37v-13.9c0-.39.29-.71.67-.76l9.35-1.21c.46-.06.87.3.87.76v6.53a5.808 5.808 0 0 0-7.71 5.49c0 3.21 2.6 5.81 5.81 5.81c3.21 0 5.81-2.6 5.81-5.81V2.869Z"/>`),
		g.Group(children),
	)
}