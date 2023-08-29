package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrophyCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTrophyCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M16.675 7.5H9.227l.635 4.418c1.222 2.623 4.956 2.623 6.178 0l.635-4.418Zm-7.448-2a2 2 0 0 0-1.98 2.285l.64 4.446c.024.17.073.336.144.492c1.925 4.217 7.915 4.217 9.84 0c.071-.156.12-.322.144-.492l.64-4.446a2 2 0 0 0-1.98-2.285H9.227Z"/><path d="M11.95 18.5v-3h2v3h-2Z"/><path d="M11.058 19.5a2.778 2.778 0 0 1 3.786 0h-3.786Zm4.91-1a4.277 4.277 0 0 0-6.033 0l-.796.791c-.632.63-.187 1.709.706 1.709h6.212c.893 0 1.338-1.08.706-1.709l-.796-.791Zm-7.547-10a1 1 0 0 1-1 1h-2.01a1 1 0 0 1 0-2h2.01a1 1 0 0 1 1 1Z"/><path d="M5.268 7.512a1.003 1.003 0 0 1 1.135.848l.07.489a3.165 3.165 0 0 0 2.62 2.667l-.328 1.968a5.168 5.168 0 0 1-4.277-4.356l-.07-.488a.997.997 0 0 1 .85-1.128Zm12.311.988a1 1 0 0 0 1 1h2.01a1 1 0 1 0 0-2h-2.01a1 1 0 0 0-1 1Z"/><path d="M20.732 7.512a1.003 1.003 0 0 0-1.135.848l-.07.489a3.166 3.166 0 0 1-2.62 2.667l.328 1.968a5.168 5.168 0 0 0 4.277-4.356l.07-.488a.997.997 0 0 0-.85-1.128Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTrophyCircleFilled0)"/></g>`),
		g.Group(children),
	)
}