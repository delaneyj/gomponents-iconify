package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hushed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.712 30.276a61.977 61.977 0 0 1-6.3-1.11a3.11 3.11 0 0 0-3 .58c-.53.54-2.06 2.05-3.65 3.61a28.69 28.69 0 0 1-13.1-13.11c1.55-1.59 3-3.11 3.59-3.64a3.11 3.11 0 0 0 .58-3a61.827 61.827 0 0 1-1.11-6.31a2 2 0 0 0-2.195-1.783a2.06 2.06 0 0 0-.105.013h-8.52a1.5 1.5 0 0 0-1.36 1.37c-.55 7.69 3.74 15.92 4.61 17.5h0v.06l.12.23h0a35.44 35.44 0 0 0 13 13h0l.44.25h0c2 1.06 9.95 5.06 17.38 4.51a1.5 1.5 0 0 0 1.39-1.36v-8.51a2 2 0 0 0-1.665-2.286a2.06 2.06 0 0 0-.105-.014Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.25 5.502c-5.113 0-9.258 3.514-9.258 7.85a7.161 7.161 0 0 0 2.234 5.105l-.982 4.465l4.218-2.412a10.627 10.627 0 0 0 3.787.692c5.113 0 9.258-3.515 9.258-7.85s-4.145-7.85-9.258-7.85Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.94 18.638c2.88-1.835 4.366-8.07-3.165-10.632c1.772 4.019-2.563 4.525-2.563 7.658c0 2.5 2.88 3.164 2.88 3.164a2.787 2.787 0 0 1-.064-2.88c.507.633 2.184 2.627 2.152-1.993a7.3 7.3 0 0 1 .76 4.683Z"/>`),
		g.Group(children),
	)
}