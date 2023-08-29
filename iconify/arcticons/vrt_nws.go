package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrtNws(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M21.186 9.218c-1.266 6.384-5.633 11.943-5.633 11.943s-3.734-5.315-5.634-11.943m14.999 4.507a4.507 4.507 0 0 1 4.507-4.507h0m-4.507 4.507v7.436"/><path d="M24.918 13.725a4.507 4.507 0 0 1 4.507-4.507h0m-4.507 4.507v7.436m8.656-8.761v4.254a4.507 4.507 0 0 0 4.507 4.507h0M33.574 5.5a4.507 4.507 0 0 0 4.507 4.507h0M8.695 42.5h30.61a2.449 2.449 0 0 0 2.449-2.449V27.807a2.449 2.449 0 0 0-2.449-2.449H8.695a2.449 2.449 0 0 0-2.449 2.45V40.05a2.449 2.449 0 0 0 2.449 2.45Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.629 30.429l-2.179 7l-2.18-7l-2.179 7l-2.179-7m10.925 6.41c.482.405 1.003.59 2.172.59h.593c.964 0 1.746-.783 1.746-1.75h0c0-.966-.782-1.75-1.746-1.75h-1.185a1.748 1.748 0 0 1-1.747-1.75h0c0-.966.782-1.75 1.746-1.75h.593c1.17 0 1.69.186 2.173.59M17.935 37.43v-4.36a2.641 2.641 0 0 0-2.641-2.64h0a2.641 2.641 0 0 0-2.642 2.64m0 4.359v-7"/>`),
		g.Group(children),
	)
}