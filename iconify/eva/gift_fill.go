package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaGiftFill0"><g id="evaGiftFill1"><path id="evaGiftFill2" fill="currentColor" d="M4.64 15.27v4.82a.92.92 0 0 0 .92.91h5.62v-5.73ZM12.82 21h5.62a.92.92 0 0 0 .92-.91v-4.82h-6.54ZM20.1 7.09h-1.84a2.82 2.82 0 0 0 .29-1.23A2.87 2.87 0 0 0 15.68 3A4.21 4.21 0 0 0 12 5.57A4.21 4.21 0 0 0 8.32 3a2.87 2.87 0 0 0-2.87 2.86a2.82 2.82 0 0 0 .29 1.23H3.9c-.5 0-.9.59-.9 1.31v3.93c0 .72.4 1.31.9 1.31h7.28V7.09h1.64v6.55h7.28c.5 0 .9-.59.9-1.31V8.4c0-.72-.4-1.31-.9-1.31Zm-11.78 0a1.23 1.23 0 1 1 0-2.45c1.4 0 2.19 1.44 2.58 2.45Zm7.36 0H13.1c.39-1 1.18-2.45 2.58-2.45a1.23 1.23 0 1 1 0 2.45Z"/></g></g>`),
		g.Group(children),
	)
}