package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashAutoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.975 8.7l-.625 1.825q-.075.2-.263.338t-.412.137q-.375 0-.588-.313T15 10.025L17.7 2.8q.125-.35.438-.55t.662-.2q.375 0 .688.238t.437.587l2.725 7.175q.125.35-.075.65T22 11q-.225 0-.425-.137T21.3 10.5l-.625-1.8h-3.7Zm.45-1.3h2.75L18.85 3.65h-.05L17.425 7.4Zm-11.9 12.925q-.225-.075-.375-.262T5 19.6V14H4q-.825 0-1.413-.588T2 12V4q0-.825.588-1.413T4 2h5.85q.8 0 1.288.625T11.425 4L10 9h1.125q.9 0 1.338.8t-.088 1.55l-6 8.675q-.15.225-.388.3t-.462 0Z"/>`),
		g.Group(children),
	)
}