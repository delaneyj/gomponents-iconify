package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.984 6.797A3.25 3.25 0 0 1 9.89 5h12.218a3.25 3.25 0 0 1 2.907 1.797l4.535 9.07c.295.59.449 1.24.449 1.9v4.983A3.25 3.25 0 0 1 26.75 26H5.25A3.25 3.25 0 0 1 2 22.75v-4.983c0-.66.154-1.31.449-1.9l4.535-9.07Zm16.244.894A1.25 1.25 0 0 0 22.108 7H9.891a1.25 1.25 0 0 0-1.118.691l-3.656 7.312C5.16 15 5.205 15 5.25 15h21.5c.045 0 .09 0 .133.003L23.228 7.69ZM4 18.25v4.5c0 .69.56 1.25 1.25 1.25h21.5c.69 0 1.25-.56 1.25-1.25v-4.5c0-.69-.56-1.25-1.25-1.25H5.25C4.56 17 4 17.56 4 18.25Zm20.5 3.5a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Z"/>`),
		g.Group(children),
	)
}