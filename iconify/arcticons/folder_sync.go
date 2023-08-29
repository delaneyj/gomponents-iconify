package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M25.251 32.965a6.585 6.585 0 0 0 3.405-11.122m-5.913-1.808a6.585 6.585 0 0 0-3.398 11.12"/><path d="m27.937 22.562l-1.75 1.75V19.5H31l-1.75 1.75l-1.313 1.312zm-7.874 7.876l1.75-1.75V33.5H17l1.75-1.75l1.313-1.312z"/></g>`),
		g.Group(children),
	)
}