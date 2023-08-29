package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loudspeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 6.07C21 4.37 22.38 3 24.07 3s3.06 1.37 3.07 3.07v6.225a3.722 3.722 0 0 1 0 7.19v6.235a3.07 3.07 0 0 1-6.14 0v-.232l-5.93-2.918a7.725 7.725 0 0 0-1.931-.648a8.166 8.166 0 0 0-1.459-.132l-.004-.01H10v6.018c0 .67-.54 1.21-1.21 1.21H6.216c-.67 0-1.21-.54-1.21-1.21V21.78h-1.28C2.772 21.78 2 21.03 2 20.11v-8.43c0-.93.77-1.68 1.727-1.68h8.045a7.844 7.844 0 0 0 3.3-.78L21 6.303V6.07ZM7.006 21.78v5.228H8V21.78h-.994ZM19.55 9.303l-3.384 1.665l-.007.004c-.992.478-2.06.798-3.159.942v7.719l.098.257c1.066.15 2.106.47 3.074.945l3.378 1.662V9.303Z"/>`),
		g.Group(children),
	)
}