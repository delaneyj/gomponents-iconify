package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M270.276 194.86H146.401v-61.905h179.17L447.701 88L242.755 318.67H366.92v61.91H187.75L65.615 424l204.661-229.14zM0 132.957l122.565-.005v61.905H55.127v123.815h82.691l-55.127 61.905H0v-247.62zm456.868 61.904h-82.69l55.131-61.905H512v247.62l-122.565.004v-61.905h67.433V194.861z"/>`),
		g.Group(children),
	)
}