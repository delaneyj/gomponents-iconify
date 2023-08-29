package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cardmindreader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.869C12.005 9.79 4.88 21.569 4.522 24.279s3.79 13.279 3.79 13.279A30.116 30.116 0 0 1 24 23.227m0 0s12.264-6.324 14.946-6.677"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.869c11.995 3.921 19.12 15.7 19.478 18.41a8.709 8.709 0 0 1-.303 2.634c-.172.766-.4 1.616-.66 2.492c-.225.766-.474 1.55-.727 2.317c-.323.977-.653 1.923-.953 2.759c-.642 1.79-1.147 3.077-1.147 3.077A30.117 30.117 0 0 0 24 23.227"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.836 34.396s-3.572-9.97-13.087-12.854M41.546 31.71a21.188 21.188 0 0 0-10.43-11.775m11.247 9.353s-3.219-7.968-8.6-10.47m9.205 7.677s-2.323-5.832-6.726-8.755M24 23.968a36.514 36.514 0 0 1-4.334 5.518c-3.249 3.602-2.947 5.923-2.947 5.923a3.233 3.233 0 0 0 2.42 3.517a2.385 2.385 0 0 0 2.669-.832l.991-1.61s.724 4.21-1.78 5.647"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.018 42.131A10.03 10.03 0 0 1 24 41.617m0-17.649a36.514 36.514 0 0 0 4.334 5.518c3.249 3.602 2.947 5.923 2.947 5.923a3.233 3.233 0 0 1-2.42 3.517a2.385 2.385 0 0 1-2.669-.832l-.991-1.61s-.724 4.21 1.78 5.647"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.982 42.131A10.03 10.03 0 0 0 24 41.617"/>`),
		g.Group(children),
	)
}