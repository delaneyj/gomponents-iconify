package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrtMax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M21.186 9.218c-1.266 6.384-5.633 11.943-5.633 11.943s-3.734-5.315-5.634-11.943m14.999 4.507a4.507 4.507 0 0 1 4.507-4.507h0m-4.507 4.507v7.436"/><path d="M24.918 13.725a4.507 4.507 0 0 1 4.507-4.507h0m-4.507 4.507v7.436m8.656-8.761v4.254a4.507 4.507 0 0 0 4.507 4.507h0M33.574 5.5a4.507 4.507 0 0 0 4.507 4.507h0M8.695 42.5h30.61a2.449 2.449 0 0 0 2.449-2.449V27.807a2.449 2.449 0 0 0-2.449-2.449H8.695a2.449 2.449 0 0 0-2.449 2.45V40.05a2.449 2.449 0 0 0 2.449 2.45Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m37.172 30.429l-5.283 7m5.283 0l-5.283-7M10.828 33.07a2.641 2.641 0 0 1 2.642-2.64h0a2.641 2.641 0 0 1 2.641 2.64v4.36m-5.283-7.001v7m5.283-4.359a2.641 2.641 0 0 1 2.642-2.64h0a2.641 2.641 0 0 1 2.642 2.64v4.36m7.812-2.642a2.641 2.641 0 0 1-2.642 2.641h0a2.641 2.641 0 0 1-2.641-2.641V33.07a2.641 2.641 0 0 1 2.641-2.642h0a2.641 2.641 0 0 1 2.642 2.642m0 4.359v-7"/>`),
		g.Group(children),
	)
}