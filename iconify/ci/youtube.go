package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Youtube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.006 19.012h-.02c-.062 0-6.265-.012-7.83-.437a2.5 2.5 0 0 1-1.764-1.765A26.494 26.494 0 0 1 1.986 12a26.646 26.646 0 0 1 .417-4.817A2.564 2.564 0 0 1 4.169 5.4c1.522-.4 7.554-.4 7.81-.4H12c.063 0 6.282.012 7.831.437c.859.233 1.53.904 1.762 1.763c.29 1.594.427 3.211.407 4.831a26.568 26.568 0 0 1-.418 4.811a2.51 2.51 0 0 1-1.767 1.763c-1.52.403-7.553.407-7.809.407Zm-2-10.007l-.005 6l5.212-3l-5.207-3Z"/>`),
		g.Group(children),
	)
}