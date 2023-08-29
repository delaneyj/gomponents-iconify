package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hatenabookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M293 390c8 6 12 16 12 31c0 13-4 22-14 28c-9 6-25 9-48 9h-33v-77h35c23 0 39 3 48 9zm-16-79c-8 5-24 7-48 7h-19v-70h21c23 0 39 3 47 8s13 15 13 28s-5 22-14 27zM80 29h490c44 0 80 36 80 80v490c0 44-36 80-80 80H80c-44 0-80-36-80-80V109c0-44 36-80 80-80zm453 154h-85v230h85V183zM384 478c8-14 12-30 12-49c0-26-7-47-21-63c-14-15-34-24-58-26c22-6 38-15 48-26c10-12 15-28 15-48c0-15-3-29-10-41s-17-22-30-29c-11-6-24-11-40-13c-16-3-43-4-83-4h-96v350h99c40 0 69-1 86-4c18-3 32-7 44-14c15-8 26-19 34-33zm139 38c9-9 13-19 13-32c0-12-4-23-13-32c-9-8-20-12-33-12s-24 4-32 13c-9 8-13 19-13 31c0 13 4 23 13 32s19 13 32 13s24-4 33-13z"/>`),
		g.Group(children),
	)
}