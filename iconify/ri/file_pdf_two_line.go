package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePdfTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4h10v4h4v12H5V4ZM3.998 2A.995.995 0 0 0 3 2.992v18.016a1 1 0 0 0 .993.992h16.014A1 1 0 0 0 21 20.992V7l-5-5H3.998ZM10.5 7.5c0 1.577-.456 3.437-1.225 5.153c-.772 1.723-1.814 3.197-2.9 4.066l1.18 1.613c2.927-1.952 6.168-3.29 9.303-2.842l.458-1.939C14.644 12.661 12.5 9.99 12.5 7.5h-2Zm.6 5.972c.267-.597.504-1.216.704-1.843a9.66 9.66 0 0 0 1.706 1.966c-.982.176-1.944.465-2.875.833c.165-.314.32-.633.465-.956Z"/>`),
		g.Group(children),
	)
}