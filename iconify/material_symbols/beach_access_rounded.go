package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeachAccessRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.9 20.3l-5.65-5.65l1.4-1.4l5.65 5.65q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275ZM4.5 18.525Q3 16.2 2.913 13.488T4.15 8.375q.075.85.425 1.913t.963 2.237Q6.15 13.7 7 14.938t1.875 2.437l-1.3 1.3q-.7.7-1.638.663T4.5 18.525ZM10.25 16q-1.2-1.2-2.1-2.613T6.738 10.65q-.513-1.325-.575-2.413t.462-1.612q.525-.55 1.613-.5t2.425.563Q12 7.2 13.413 8.113t2.612 2.112L10.25 16Zm8.425-8.425L17.4 8.85Q16.225 7.825 14.988 7t-2.4-1.45q-1.163-.625-2.225-.988T8.45 4.125Q10.825 2.9 13.525 3T18.5 4.525q.775.5.825 1.425t-.65 1.625Z"/>`),
		g.Group(children),
	)
}