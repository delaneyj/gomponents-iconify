package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonWithCrown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="m31.44 14.337l.002-9.743l3.115 2.652L35.972 3.7l1.469 3.409l2.844-2.515v9.743H31.44z"/><path fill="#92d3f5" d="M17 61v-4c0-4.994 5.008-9 10-9c6 5 12 5 18 0c4.994 0 10 4.006 10 9v4"/><path fill="#fcea2b" d="M24.936 31c0 9 4.937 14 11 14C41.873 45 47 40 47 31a12.137 12.137 0 0 0-1-5c-3-3-7-8-7-8c-4 3-7 6-13 7c0 0-1.064 1-1.064 6Z"/><path fill="#f1b31c" d="M31.44 11.585C24.764 13.45 22 19.485 22 25.518C22 32.779 22 39 26 39l-.785-5.786l.552-7.739l8.43-4.007l5.252-2.971L46 26l.977 6L46 39c4 0 4-6.222 4-13.482c0-6.116-2.84-12.232-9.715-14.007v2.826H31.44v-2.752Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M44.097 14.089C48.23 16.684 50 21.342 50 26c0 7 0 13-4 13m-20 0c-4 0-4-6-4-13c0-4.657 1.77-9.315 5.901-11.91M17 60v-3c0-4.994 5.008-9 10-9c6 5 12 5 18 0c4.994 0 10 4.006 10 9v3"/><path d="M41.873 30a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-8 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2" d="M24.936 31c0 9 4.937 14 11 14C41.873 45 47 40 47 31a12.137 12.137 0 0 0-1-5c-3-3-7-8-7-8c-4 3-7 6-13 7c0 0-1.064 1-1.064 6Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M33 38a6.55 6.55 0 0 0 6 0m-7.56-23.663l.002-9.743l3.115 2.652L35.972 3.7l1.469 3.409l2.844-2.515v9.743H31.44z"/>`),
		g.Group(children),
	)
}