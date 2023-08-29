package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emphasis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M525 594s34 4 35 4c-4 43-27 110-40 132c-51-1-143-2-222-2H173c-46 0-101 1-173 2v-29l53-5c44-6 49-11 49-87V301c0-75-5-80-49-88l-30-4v-29c45 2 101 3 147 3h203c62 0 116-1 129-3c1 13 7 72 12 127l-36 4c-10-29-23-53-36-65c-18-15-47-25-94-25h-77c-31 0-31 1-31 32v144c0 24 1 25 27 25h66c50 0 61-6 71-45l5-21s35 0 35 1c-2 26-3 56-3 87c0 32 1 62 3 88c0 1-36 1-35 1l-5-21c-10-39-21-46-71-46h-66c-26 0-27 1-27 26v98c0 38 4 64 15 78c13 15 31 22 101 22c83 0 116-3 169-96zM340 70c0-39-31-70-70-70c-38 0-70 31-70 70c0 38 32 70 70 70c39 0 70-32 70-70z"/>`),
		g.Group(children),
	)
}