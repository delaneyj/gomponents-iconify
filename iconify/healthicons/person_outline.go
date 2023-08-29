package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 24c4.42 0 8-3.58 8-8s-3.58-8-8-8s-8 3.58-8 8s3.58 8 8 8Zm10-8c0 5.525-4.475 10-10 10s-10-4.475-10-10S18.475 6 24 6s10 4.475 10 10ZM9.223 34.212C8.22 35.022 8 35.629 8 36v4h32v-4c0-.37-.22-.979-1.224-1.788c-.98-.791-2.442-1.545-4.214-2.197C31.02 30.712 26.753 30 24 30c-2.753 0-7.02.712-10.562 2.015c-1.772.652-3.234 1.406-4.215 2.197ZM24 28c-6.008 0-18 3.035-18 8v6h36v-6c0-4.965-11.992-8-18-8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}