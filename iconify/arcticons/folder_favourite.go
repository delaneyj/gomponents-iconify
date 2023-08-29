package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderFavourite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 19.507l1.652 5.086H31l-4.326 3.143l1.652 5.086L24 29.678l-4.326 3.144l1.652-5.086L17 24.593h5.348L24 19.507z"/>`),
		g.Group(children),
	)
}