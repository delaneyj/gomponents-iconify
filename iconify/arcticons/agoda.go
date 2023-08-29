package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Agoda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" class="d"><path d="M18.992 19.322v5.822a1.947 1.947 0 0 1-1.942 1.941c-.582 0-1.068-.194-1.36-.582"/><path d="M17.05 19.322c1.068 0 1.942.874 1.942 1.941v1.261c0 1.068-.874 1.941-1.942 1.941s-1.942-.873-1.942-1.94v-1.262c0-1.067.874-1.94 1.942-1.94Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 24.466a1.947 1.947 0 0 1-1.942-1.941v-1.262c0-1.067.874-1.94 1.942-1.94s1.942.873 1.942 1.94v1.262c0 1.067-.874 1.94-1.942 1.94Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.043 22.524c0 1.068-.874 1.941-1.942 1.941s-1.943-.873-1.943-1.94v-1.262c0-1.067.874-1.94 1.943-1.94s1.942.873 1.942 1.94m0 3.202v-5.143m20.849 1.941c0-1.067-.874-1.94-1.942-1.94s-1.942.873-1.942 1.94v1.262c0 1.067.874 1.94 1.942 1.94s1.942-.873 1.942-1.94m0 1.94v-7.763m6.95 5.822c0 1.068-.874 1.941-1.942 1.941s-1.943-.873-1.943-1.94v-1.262c0-1.067.874-1.94 1.943-1.94s1.942.873 1.942 1.94m0 3.202v-5.143" class="d"/><circle cx="10.101" cy="30.52" r=".751" fill="currentColor"/><circle cx="17.05" cy="30.52" r=".751" fill="currentColor"/><circle cx="24" cy="30.52" r=".751" fill="currentColor"/><circle cx="30.95" cy="30.52" r=".751" fill="currentColor"/><circle cx="37.9" cy="30.548" r=".751" fill="currentColor"/>`),
		g.Group(children),
	)
}