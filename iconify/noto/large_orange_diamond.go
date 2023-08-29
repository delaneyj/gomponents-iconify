package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LargeOrangeDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#F77E00" d="m70.04 6.32l52.2 52a7.887 7.887 0 0 1-.2 11.1l-52 52.2a7.991 7.991 0 0 1-11.3 0l-52-52.2a8 8 0 0 1-.2-11.1l52.2-52a7.991 7.991 0 0 1 11.3 0z"/><path fill="#FF9800" d="m65.14 10.42l48.5 48.4a7.268 7.268 0 0 1-.18 10.28l-.02.02l-48.4 48.5a7.42 7.42 0 0 1-10.49.01l-.01-.01l-48.3-48.5a7.28 7.28 0 0 1-.21-10.29l.01-.01l48.6-48.4c2.9-2.9 7.59-2.9 10.5 0z"/><path fill="#FFBD52" d="M53.34 17.62c-1.6 1.1-33.1 32.3-33.1 32.3s-2.6 2.3-.8 4.7c1.7 2.2 4.3 1.1 5.6-.1s29.7-29 30.9-30.4a5.425 5.425 0 0 0 .6-5.5c-.8-1.6-2.1-1.8-3.2-1z"/><circle cx="15.94" cy="60.43" r="3.3" fill="#FFBD52"/>`),
		g.Group(children),
	)
}